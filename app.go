package main

import (
	"git-uwu/actions"
	"git-uwu/uwuify"
	"github.com/urfave/cli/v2"
	"log"
	"os"
)

func main() {
	app := &cli.App{
		Commands: []*cli.Command{
			{
				Name:        "commit",
				Aliases:     []string{"c"},
				Description: uwuify.From("Adds a commit to Git."),
				Action:      actions.Commit,
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:    "message",
						Aliases: []string{"m"},
						Usage:   uwuify.From("Adds a commit message to the commit."),
					},
				},
			},
			{
				Name:        "add",
				Aliases:     []string{"a"},
				Description: uwuify.From("Adds one or more files to be tracked."),
				Action:      actions.Add,
				Flags: []cli.Flag{
					&cli.PathFlag{
						Name:     "path",
						Usage:    uwuify.From("The path to commit to Git."),
						Required: true,
						Hidden:   true,
					},
				},
			},
			{
				Name:        "magic",
				Description: uwuify.From("Does some magic on a piece of text."),
				Action:      actions.Magic,
				Flags: []cli.Flag{
					&cli.PathFlag{
						Name:     "file",
						Usage:    uwuify.From("The file to turn into magic."),
						Required: false,
					},
					&cli.StringFlag{
						Name:     "message",
						Aliases:  []string{"m"},
						Usage:    uwuify.From("Adds magic to a message."),
						Required: false,
					},
				},
				Hidden: true,
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
