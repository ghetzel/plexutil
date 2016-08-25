package main

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"github.com/fatih/color"
	"github.com/ghetzel/cli"
	"github.com/ghetzel/plexutil"
	"github.com/ghetzel/plexutil/client"
	"github.com/ghodss/yaml"
	"github.com/op/go-logging"
	"io"
	"math"
	"os"
	"text/tabwriter"
)

var log = logging.MustGetLogger(`main`)
var plex *client.PlexClient

func main() {
	app := cli.NewApp()
	app.Name = `plexutil`
	app.Usage = `Manage a Plex Media Server from the command line`
	app.Version = `0.0.1`
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
			Value: `basic`,
		},
		cli.StringFlag{
			Name:  `url, U`,
			Usage: `The base URL used to access the Plex Media Server`,
		},
	}

	app.Before = func(c *cli.Context) error {
		if config, err := plexutil.LoadConfig(c.String(`config`)); err == nil || os.IsNotExist(err) {
			if url := c.String(`url`); url != `` {
				config.URL = url
			}

			plex = client.NewFromConfig(config)
		} else {
			log.Fatal(err)
		}

		return nil
	}

	app.Commands = []cli.Command{
		{
			Name:      `call`,
			Usage:     `Perform a generic API call against the configured Plex Media Server`,
			ArgsUsage: ``,
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
					Name:  `option, o`,
					Usage: `An query string value to include with the request (specified as key=value)`,
				},
			},
			Action: func(c *cli.Context) {
				log.Fatalf("NOT IMPLEMENTED")
			},
		},
		{
			Name:  `sessions`,
			Usage: `List all currently playing media`,
			Action: func(c *cli.Context) {
				if videos, err := plex.CurrentSessions(); err == nil {
					printWithFormat(c.GlobalString(`format`), videos, func() {
						tw := tabwriter.NewWriter(os.Stdout, 0, 0, 1, ' ', 0)

						for _, video := range videos {

							fmt.Fprintf(tw, "%s\t%s\t",
								video.User.Title,
								video.Player.State)

							printAsciiProgressBar(tw, video.TranscodeSession.Progress, 20, video.Player.State)

							fmt.Fprintf(tw, "\t%s\tS%02dE%02d\t%s\t%s\t%s\n",
								video.GrandparentTitle,
								video.ParentIndex,
								video.Index,
								video.Title,
								video.Player.Address,
								video.Player.Title)
						}

						tw.Flush()
					})
				} else {
					log.Fatal(err)
				}
			},
		},
	}

	app.Run(os.Args)
}

func printWithFormat(format string, data interface{}, fallbackFunc func()) {
	var output []byte
	var err error

	switch format {
	case `json`:
		output, err = json.Marshal(data)
	case `yaml`:
		output, err = yaml.Marshal(data)
	case `xml`:
		output, err = xml.Marshal(data)
	default:
		fallbackFunc()
		return
	}

	if err == nil {
		fmt.Println(string(output[:]))
	} else {
		log.Fatal(err)
	}
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
