package main

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"io"
	"math"
	"os"
	"regexp"
	"strings"
	"text/tabwriter"
	"time"

	"github.com/ghetzel/cli"
	"github.com/ghetzel/go-stockutil/httputil"
	"github.com/ghetzel/go-stockutil/log"
	"github.com/ghetzel/go-stockutil/stringutil"
	"github.com/ghetzel/go-stockutil/typeutil"
	"github.com/ghetzel/plexutil"
	"github.com/ghetzel/plexutil/client"
	"github.com/ghodss/yaml"
)

const DEFAULT_FORMAT = `basic`

var normRx = regexp.MustCompile(`(?:[\s\W\-\_]+)`)
var plex *client.PlexClient

func main() {
	app := cli.NewApp()
	app.Name = plexutil.Name
	app.Usage = plexutil.Description
	app.Version = plexutil.Version
	app.EnableBashCompletion = false

	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:   `log-level, L`,
			Usage:  `Level of log output verbosity`,
			Value:  `info`,
			EnvVar: `LOGLEVEL`,
		},
		cli.StringFlag{
			Name:  `config, c`,
			Usage: `The path to the configuration file`,
			Value: `~/.config/plexutil/config.yml`,
		},
		cli.StringFlag{
			Name:  `format, f`,
			Usage: `The output format to use when printing results`,
			Value: DEFAULT_FORMAT,
		},
		cli.StringFlag{
			Name:   `url, U`,
			Usage:  `The base URL used to access the Plex Media Server`,
			EnvVar: `PLEXUTIL_URL`,
		},
		cli.StringFlag{
			Name:   `token, K`,
			Usage:  `Specify the token use when authorizing against the Plex Media Server.`,
			EnvVar: `PLEXUTIL_TOKEN`,
		},
		cli.BoolFlag{
			Name:  `ignore-ssl-verify, S`,
			Usage: `Whether to ignore SSL verification.`,
		},
		cli.DurationFlag{
			Name:  `timeout, T`,
			Usage: `HTTP client timeout for communicating with the Plex server.`,
			Value: 5 * time.Second,
		},
	}

	app.Before = func(c *cli.Context) error {
		log.SetLevelString(c.String(`log-level`))

		if config, err := plexutil.LoadConfig(c.String(`config`)); err == nil || os.IsNotExist(err) {
			if url := c.String(`url`); url != `` {
				config.URL = url
			}

			if config.URL == `` {
				log.Fatalf("Must specify a Plex Media Serve URL, via --url/-u or the PLEXUTIL_URL envvar.")
			}

			plex = client.NewFromConfig(config)
			plex.IgnoreSSL = c.Bool(`ignore-ssl-verify`)
			plex.Timeout = c.Duration(`timeout`)
			plex.SetAuthToken(c.String(`token`))
		} else {
			log.Fatalf("config load error: %v", err)
		}

		return nil
	}

	app.Commands = []cli.Command{
		{
			Name:      `call`,
			Usage:     `Perform a generic API call against the configured Plex Media Server`,
			ArgsUsage: `PATH`,
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  `method, m`,
					Usage: `The HTTP method to use`,
					Value: `get`,
				},
				cli.StringSliceFlag{
					Name:  `header, H`,
					Usage: `An HTTP header to include with the request (specified as 'Header Name=Header Value')`,
				},
				cli.StringSliceFlag{
					Name:  `query, q`,
					Usage: `An query string value to include with the request (specified as key=value)`,
				},
			},
			Action: func(c *cli.Context) {
				var body interface{}
				var responseBody client.MediaContainer
				var params = make(map[string]interface{})
				var header = make(map[string]interface{})

				for _, qs := range c.StringSlice(`query`) {
					var k, v = stringutil.SplitPair(qs, `=`)
					params[k] = typeutil.Auto(v)
				}

				for _, qs := range c.StringSlice(`header`) {
					var k, v = stringutil.SplitPair(qs, `=`)
					params[k] = typeutil.Auto(v)
				}

				if res, err := plex.Request(
					httputil.Method(c.String(`method`)),
					c.Args().First(),
					body,
					params,
					header,
				); err == nil {
					if err := plex.Decode(res.Body, &responseBody); err == nil {
						var format = c.GlobalString(`format`)

						if format == DEFAULT_FORMAT {
							format = `yaml`
						}

						printWithFormat(format, responseBody)
					} else {
						log.Fatalf("decode error: %v", err)
					}
				} else {
					log.Fatalf("request error: %v", err)
				}
			},
		},
		{
			Name:  `sessions`,
			Usage: `List all currently playing media`,
			Flags: []cli.Flag{
				cli.BoolFlag{
					Name:  `history, H`,
					Usage: `Whether to display all session history or only active sessions`,
				},
				cli.IntFlag{
					Name:  `per-page, l`,
					Usage: `Number of results to query at a time`,
					Value: 25,
				},
				cli.IntFlag{
					Name:  `page, p`,
					Usage: `The page number to show`,
					Value: 1,
				},
			},
			Action: func(c *cli.Context) {
				var videos []client.Video

				if c.Bool(`history`) {
					if v, err := plex.RecentSessions(c.Int(`per-page`), c.Int(`page`)); err == nil {
						videos = v
					} else {
						log.Fatal(err)
						return
					}
				} else {
					if v, err := plex.CurrentSessions(); err == nil {
						videos = v
					} else {
						log.Fatal(err)
						return
					}
				}

				printWithFormat(c.GlobalString(`format`), videos, func() {
					if len(videos) == 0 {
						return
					}

					tw := tabwriter.NewWriter(os.Stdout, 0, 0, 1, ' ', 0)

					fmt.Fprintf(tw, "User\tState\tProgress\tShow\tEpisode\tTitle\tAddress\tDevice\n")

					for _, video := range videos {
						state := video.Player.State

						switch state {
						case `paused`:
							state = `${yellow}` + state + `${reset}`
						case `playing`:
							state = `${green}` + state + `${reset}`
						case `buffering`:
							state = `${red}` + state + `${reset}`
						}

						log.CFPrintf(tw, "%s\t%s\t", video.User.Title, state)

						if video.Player.State == `` {
							if video.ViewedAt > 0 {
								fmt.Fprintf(tw, "%s", time.Unix(video.ViewedAt, 0).Format(`2006-01-02 15:04-0700`))
							}
						} else {
							percent := (float64(video.ViewOffset) / float64(video.Duration)) * 100.0

							printAsciiProgressBar(tw, percent, 20, video.Player.State)
							fmt.Fprintf(
								tw,
								" %2.1f%% %s/%s",
								percent,
								video.TimeElapsed(),
								video.TranscodeSession.TimeDuration(),
							)
						}

						fmt.Fprintf(tw, "\t%s\t%s\t%s\t%s\t%s\n",
							video.GrandparentTitle,
							getEpisodeNumber(&video),
							video.Title,
							video.Player.Address,
							video.Player.Title)
					}

					tw.Flush()
				})
			},
		},
		{
			Name:  `ls`,
			Usage: `List all objects in a given path`,
			Flags: []cli.Flag{
				cli.BoolFlag{
					Name:  `id, I`,
					Usage: `Only list the URL paths for accessing child objects`,
				},
				cli.StringFlag{
					Name:  `area, a`,
					Usage: `Which media area to enumerate`,
					Value: `library`,
				},
			},
			Action: func(c *cli.Context) {
				area := c.String(`area`)
				path := strings.Split(strings.Trim(c.Args().First(), `/`), `/`)

				if results, err := plex.ListDirectory(area, path); err == nil {
					printWithFormat(c.GlobalString(`format`), results, func() {
						tw := tabwriter.NewWriter(os.Stdout, 0, 0, 1, ' ', 0)

						if c.Bool(`id`) {
							for _, directory := range results.Directories {
								fmt.Fprintf(tw, "%s\n", directory.PathKey)
							}

							for _, video := range results.Videos {
								fmt.Fprintf(tw, "%d\n", video.RatingKey)
							}
						} else {

							if len(results.Directories) > 0 {
								fmt.Fprintf(tw, "Subdirectories:\tPath Key\tFriendly Name\tSection\n")

								// list directories
								for _, directory := range results.Directories {
									fmt.Fprintf(tw, "d---\t%s\t%s\t%s\n",
										directory.PathKey,
										directory.Title,
										results.LibrarySectionTitle)
								}
							}

							if len(results.Videos) > 0 {
								fmt.Fprintf(tw, "Videos:\tID\tType\tShow\tEpisode\tTitle\tDuration\tAired At\n")

								// list videos
								for _, video := range results.Videos {
									fmt.Fprintf(tw, "f---\t%d\t%s\t%s\t%s\t%s\t%s\t%s\n",
										video.RatingKey,
										video.Type,
										video.GrandparentTitle,
										getEpisodeNumber(&video),
										video.Title,
										getReadableTime(video.Duration),
										video.OriginallyAvailableAt)
								}
							}

						}

						tw.Flush()
					})
				} else {
					log.Fatalf("list error: %v", err)
					return
				}
			},
		}, {
			Name:  `info`,
			Usage: `Show information about a specific media entry`,
			Flags: []cli.Flag{
				cli.BoolFlag{
					Name:  `source-file, F`,
					Usage: `Only output the source file(s) of the media file.`,
				},
				cli.BoolFlag{
					Name:  `url, U`,
					Usage: `Only output the Plex URL(s) of the media file.`,
				},
			},
			Action: func(c *cli.Context) {
				if id, err := stringutil.ConvertToInteger(c.Args().First()); err == nil {
					if video, err := plex.GetMetadata(int(id)); err == nil {
						printWithFormat(c.GlobalString(`format`), video, func() {
							if c.Bool(`source-file`) {
								for _, part := range video.Media.Parts {
									fmt.Println(part.File)
								}
							} else if c.Bool(`url`) {
								for _, part := range video.Media.Parts {
									fmt.Printf("%s%s\n", plex.Address(), part.Key)
								}
							} else {
								tw := tabwriter.NewWriter(os.Stdout, 0, 0, 1, ' ', 0)

								fmt.Fprintf(tw, "Section:\t%s\n", video.LibrarySectionTitle)
								fmt.Fprintf(tw, "Show:\t%s\n", video.GrandparentTitle)
								fmt.Fprintf(tw, "Title:\t%s\n", video.Title)
								fmt.Fprintf(tw, "Season:\t%d\n", video.ParentIndex)
								fmt.Fprintf(tw, "Episode:\t%d\n", video.Index)
								fmt.Fprintf(tw, "Duration:\t%s\n", getReadableTime(video.Duration))
								fmt.Fprintf(tw, "Summary:\t%s\n", video.Summary)
								fmt.Fprintf(tw, "Rating:\t%s\n", video.ContentRating)
								fmt.Fprintf(tw, "Resolution:\t%dx%v\n", video.Media.Width, video.Media.VideoResolution)
								fmt.Fprintf(tw, "Framerate:\t%s\n", video.Media.VideoFrameRate)
								fmt.Fprintf(tw, "Parts:\n")

								if len(video.Media.Parts) > 0 {
									for i, part := range video.Media.Parts {
										fmt.Fprintf(tw, "%d:\n", i)
										fmt.Fprintf(tw, "  ID:\t%d\n", part.Id)
										fmt.Fprintf(tw, "  URL:\t%s%s\n", plex.Address(), part.Key)
										fmt.Fprintf(tw, "  SourceFile:\t%s\n", part.File)
										fmt.Fprintf(tw, "  Duration:\t%s\n", getReadableTime(part.Duration))

										if v, err := stringutil.ToByteString(part.Size, `%.2f`); err == nil {
											fmt.Fprintf(tw, "  Size:\t%s\n", v)
										}

										fmt.Fprintf(tw, "  Container:\t%s\n", part.Container)
										fmt.Fprintf(tw, "  Streams:\n")

										for j, stream := range part.Streams {
											fmt.Fprintf(tw, "  %d:\n", j)

											switch stream.StreamType {
											case 1:
												fmt.Fprintf(tw, "    Type:\tvideo\n")
												fmt.Fprintf(tw, "    Resolution:\t%dx%d\n", stream.Width, stream.Height)
												fmt.Fprintf(tw, "    Codec:\t%s (%s)\n", stream.Codec, stream.CodecID)
												fmt.Fprintf(tw, "    Duration:\t%s\n", getReadableTime(stream.Duration))
												fmt.Fprintf(tw, "    Framerate:\t%g (%s)\n", stream.FrameRate, stream.FrameRateMode)
												fmt.Fprintf(tw, "    BPP:\t%d\n", stream.BitDepth)
												fmt.Fprintf(tw, "    PixelFormat:\t%s\n", stream.PixelFormat)

											case 2:
												fmt.Fprintf(tw, "    Type:\taudio\n")
												fmt.Fprintf(tw, "    Codec:\t%s (%s)\n", stream.Codec, stream.CodecID)
												fmt.Fprintf(tw, "    Duration:\t%s\n", getReadableTime(stream.Duration))
												fmt.Fprintf(tw, "    Bitrate:\t%dKbps\n", stream.Bitrate)
												fmt.Fprintf(tw, "    BitsPerSample:\t%d\n", stream.BitDepth)

											case 3:
												fmt.Fprintf(tw, "    Type:\tsubtitles\n")
												fmt.Fprintf(tw, "    Codec:\t%s (%s)\n", stream.Codec, stream.CodecID)
											}

											if stream.Profile != `` {
												fmt.Fprintf(tw, "    Profile:\t%s\n", stream.Profile)
											}

											fmt.Fprintf(tw, "    StreamID:\t%d\n", stream.Id)
										}

										fmt.Fprintf(tw, "\n")
									}
								}

								tw.Flush()
							}
						})
					} else {
						log.Fatal(err)
					}
				} else {
					log.Fatal(err)
				}
			},
		},
	}

	app.Run(os.Args)
}

func printWithFormat(format string, data interface{}, fallbackFunc ...func()) {
	var output []byte
	var err error

	switch format {
	case `json`:
		output, err = json.MarshalIndent(data, ``, `  `)
	case `yaml`:
		output, err = yaml.Marshal(data)
	case `xml`:
		output, err = xml.MarshalIndent(data, ``, `  `)
	default:
		if len(fallbackFunc) > 0 {
			f := fallbackFunc[0]
			f()
			return
		} else if data != nil {
			if v, err := stringutil.ToString(data); err == nil {
				output = []byte(v[:])
			} else {
				log.Fatalf("format error: %v", err)
			}
		}
	}

	if err == nil {
		fmt.Println(string(output[:]))
	} else {
		log.Fatalf("format error: %v", err)
	}
}

func getReadableTime(msec int) string {
	sec := int64(float64(msec) / float64(1000))
	return strings.TrimLeft(time.Unix(sec, 0).UTC().Format("15:04:05"), `0:`)
}

func getEpisodeNumber(video *client.Video) string {
	if video != nil {
		if video.GrandparentTitle != `` {
			return fmt.Sprintf("S%02dE%02d", video.ParentIndex, video.Index)
		}
	}

	return ``
}

func normalizeEntry(input string) string {
	input = normRx.ReplaceAllString(input, ``)
	input = strings.ToLower(input)

	return input
}

func entryMatches(have string, want string) bool {
	if strings.Contains(normalizeEntry(have), normalizeEntry(want)) {
		return true
	}

	return false
}

func printAsciiProgressBar(w io.Writer, percent float64, totalChars int, state string) {
	output := ``
	numOn := int(math.Ceil(math.Mod(percent, float64(totalChars))))
	numOff := (totalChars - numOn)

	for i := 0; i < numOn; i++ {
		output = output + fmt.Sprintf("%c", 0x2593)
	}

	for i := 0; i < numOff; i++ {
		output = output + fmt.Sprintf("%c", 0x2591)
	}

	fmt.Fprintf(w, output)
}
