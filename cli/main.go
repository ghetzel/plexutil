package main

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"github.com/fatih/color"
	"github.com/ghetzel/bee-hotel"
	"github.com/ghetzel/cli"
	"github.com/ghetzel/go-stockutil/stringutil"
	"github.com/ghetzel/plexutil"
	"github.com/ghetzel/plexutil/client"
	"github.com/ghodss/yaml"
	"github.com/op/go-logging"
	"io"
	"math"
	"os"
	"strings"
	"text/tabwriter"
	"time"
)

const DEFAULT_FORMAT = `basic`

var log = logging.MustGetLogger(`main`)
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
			Name:  `url, U`,
			Usage: `The base URL used to access the Plex Media Server`,
		},
	}

	app.Before = func(c *cli.Context) error {
		logging.SetFormatter(logging.MustStringFormatter(`%{color}%{level:.4s}%{color:reset}[%{id:04d}] %{message}`))

		if level, err := logging.LogLevel(c.String(`log-level`)); err == nil {
			logging.SetLevel(level, `main`)
			logging.SetLevel(level, `plex`)
		} else {
			return err
		}

		if config, err := plexutil.LoadConfig(c.String(`config`)); err == nil || os.IsNotExist(err) {
			if url := c.String(`url`); url != `` {
				config.URL = url
			}

			plex = client.NewFromConfig(config)

			if err := plex.Initialize(); err != nil {
				log.Fatal(err)
			}
		} else {
			log.Fatal(err)
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

				if _, err := plex.Request(strings.ToUpper(c.String(`method`)), c.Args().First(), body, &responseBody, nil, func(request *bee.MultiClientRequest) error {
					for _, qs := range c.StringSlice(`query`) {
						pair := strings.SplitN(qs, `=`, 2)
						request.QuerySet(pair[0], pair[1])
					}

					for _, header := range c.StringSlice(`header`) {
						pair := strings.SplitN(header, `=`, 2)
						request.HeaderSet(pair[0], pair[1])
					}

					return nil

				}); err == nil {
					format := c.GlobalString(`format`)

					if format == DEFAULT_FORMAT {
						format = `yaml`
					}

					printWithFormat(format, responseBody)
				} else {
					log.Fatal(err)
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
					tw := tabwriter.NewWriter(os.Stdout, 0, 0, 1, ' ', 0)

					fmt.Fprintf(tw, "User\tState\tProgress\tShow\tEpisode\tTitle\tAddress\tDevice\n")

					for _, video := range videos {

						fmt.Fprintf(tw, "%s\t%s\t",
							video.User.Title,
							video.Player.State)

						if video.Player.State == `` {
							if video.ViewedAt > 0 {
								fmt.Fprintf(tw, "%s", time.Unix(video.ViewedAt, 0).Format(`2006-01-02 15:04-0700`))
							}
						} else {
							printAsciiProgressBar(tw, video.TranscodeSession.Progress, 20, video.Player.State)
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
		}, {
			Name:  `ls`,
			Usage: `List all objects in a given path`,
			Flags: []cli.Flag{
				cli.BoolFlag{
					Name:  `id`,
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

						if len(results.Directories) > 0 {
							fmt.Fprintf(tw, "Subdirectories:\tPath Key\tFriendly Name\tSection\n")

							// list directories
							for _, directory := range results.Directories {
								parts := strings.SplitN(directory.PathKey, `?`, 2)
								key := parts[0]

								fmt.Fprintf(tw, "d---\t%s\t%s\t%s\n",
									key,
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

						tw.Flush()
					})
				} else {
					log.Fatal(err)
					return
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
				log.Fatal(err)
			}
		}
	}

	if err == nil {
		fmt.Println(string(output[:]))
	} else {
		log.Fatal(err)
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

func printAsciiProgressBar(w io.Writer, percent float64, totalChars int, state string) {
	color.Output = w

	output := ``
	numOn := int(math.Ceil(math.Mod(percent, float64(totalChars))))
	numOff := (totalChars - numOn)

	for i := 0; i < numOn; i++ {
		output = output + fmt.Sprintf("%c", 0x2593)
	}

	for i := 0; i < numOff; i++ {
		output = output + fmt.Sprintf("%c", 0x2591)
	}

	switch state {
	case `playing`:
		color.New(color.FgGreen).Printf("%s", output)
	case `paused`:
		color.New(color.FgRed).Printf("%s", output)
	default:
		fmt.Fprintf(w, output)
	}
}
