package main // import "github.com/bbiskup/edify"

import (
	"errors"
	"github.com/bbiskup/edify/commands"
	"github.com/codegangsta/cli"
	"log"
	"os"
	"time"
)

var (
	versionFlag = cli.StringFlag{
		Name:  "version",
		Value: "14B",
		Usage: "UNCE EDIFACT version (e.g. '14B')",
	}
	specDirFlag = cli.StringFlag{
		Name:  "specdir, d",
		Value: "",
		Usage: "Specification directory (UNCE layout)",
	}
)

func main() {
	app := cli.NewApp()
	app.Name = "edify"
	app.Usage = "EDIFACT tool"
	app.EnableBashCompletion = true

	var err error

	app.Commands = []cli.Command{
		{
			Name:    "download_specs",
			Usage:   "Download specs from remote server",
			Aliases: []string{"d"},
			Action: func(c *cli.Context) {
				err = commands.DownloadSpecs(c.String("version"))
			},
			Flags: []cli.Flag{
				versionFlag,
			},
		},
		{
			Name:    "extract_specs",
			Usage:   "Extracts previously downloaded specs",
			Aliases: []string{"x"},
			Action: func(c *cli.Context) {
				err = commands.ExtractSpecs(c.String("version"))
			},
			Flags: []cli.Flag{
				versionFlag,
			},
		},
		{
			Name:    "purge_specs",
			Usage:   "Purge previously extracted specs",
			Aliases: []string{"u"},
			Action: func(c *cli.Context) {
				purgeAll := c.Bool("all")
				err = commands.PurgeSpecs(c.String("version"), purgeAll)
			},

			Flags: []cli.Flag{
				versionFlag,
				cli.BoolFlag{
					Name:  "all, a",
					Usage: "delete everything (including downloaded archives)",
				},
			},
		},
		{
			Name:    "parse",
			Usage:   "Parse a particular spec file",
			Aliases: []string{"p"},
			Action: func(c *cli.Context) {
				fileNames := c.Args()
				err = commands.Parse(fileNames)
			},
		},

		{
			Name:    "query",
			Usage:   "Query a message part",
			Aliases: []string{"q"},
			Action: func(c *cli.Context) {
				if len(c.Args()) > 0 {
					err = errors.New("Unexpected arguments")
					return
				}

				msg := c.String("msg")

				err = commands.Query(
					c.String("version"), c.String("specdir"),
					msg,
					// dump only makes sense if a message is specified
					c.Bool("dump-message") && msg != "",
					c.String("query"))
			},
			Flags: []cli.Flag{
				versionFlag,
				specDirFlag,
				cli.StringFlag{
					Name:  "msg, m",
					Value: "",
					Usage: "EDIFACT message file",
				},
				cli.StringFlag{
					Name:  "dump-message, D",
					Value: "true",
					Usage: "Dump parsed EDIFACT message",
				},
				cli.StringFlag{
					Name:  "query, q",
					Value: "",
					Usage: "Query string",
				},
			},
		},
	}

	start := time.Now()

	app.Run(os.Args)
	if err != nil {
		log.Printf("Error: %s\n", err)
		os.Exit(1)
	}

	end := time.Now()
	duration := end.Sub(start)

	log.Printf("Duration: %d ms", duration.Nanoseconds()/1e6)

}
