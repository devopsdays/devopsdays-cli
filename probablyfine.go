package main

import (
	"fmt"
	"os"

	"gopkg.in/urfave/cli.v1"
)

const version = "0.0.1"

func main() {
	app := cli.NewApp()
	app.Version = version
	app.Name = "probablyfine"
	app.Usage = "Run maintainence tasks for the devopsdays.org website"
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "city, c",
			Value: "Chicago",
			Usage: "The city for the event. If there are spaces, surround it with quotes",
		},
	}
	app.Commands = []cli.Command{
		{
			Name:    "event",
			Aliases: []string{"e"},
			Usage:   "options for events",
			Subcommands: []cli.Command{
				{
					Name:  "add",
					Usage: "add a new event",
					Action: func(c *cli.Context) error {
						city := c.GlobalString("city")
						fmt.Printf("new event for %s added\n", city)
						return nil
					},
				},
			},
		},
		{
			Name:    "sponsor",
			Aliases: []string{"s"},
			Usage:   "options for sponsors",
			Subcommands: []cli.Command{
				{
					Name:  "add",
					Usage: "add a new sponsor",
					Action: func(c *cli.Context) error {
						sponsor := "Chef"
						fmt.Printf("new sponsor for %s added\n", sponsor)
						return nil
					},
				},
				{
					Name:  "audit",
					Usage: "audit all sponsors for logos and proper size",
					Action: func(c *cli.Context) error {
						logmsg := "All sponsors look fine"
						fmt.Printf("%s\n", logmsg)
						return nil
					},
				},
			},
		},
	}

	app.Run(os.Args)
}

func addEvent() { // TODO: write add event function

}

func addSponsor() { // TODO: write add sponsor function

}
