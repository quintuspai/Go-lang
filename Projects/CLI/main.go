package main

import (
	"fmt"
	"log"
	"os"

	"github.com/urfave/cli"
)

func main() {

	app := &cli.App{
		Name:                   "boom",
		Usage:                  "make an explosion",
		Version:                "1.0.0",
		UseShortOptionHandling: true,
		EnableBashCompletion:   true,
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:  "lang",
				Value: "english",
				Usage: "language for greeting",
			},
		},

		Action: func(c *cli.Context) error {
			name := "NA"
			if c.NArg() > 0 {
				name = c.Args().Get(0)
			}
			if c.String("lang") == "spanish" {
				fmt.Println("Hola", name)
			} else {
				fmt.Println("Hello", name)
			}
			return nil
		},

		Commands: []cli.Command{
			{
				Name:    "add",
				Aliases: []string{"a"},
				Usage:   "add a task to the list",
				Action: func(c *cli.Context) error {
					fmt.Println("added task: ", c.Args().First())
					return nil
				},
			},
			{
				Name:    "complete",
				Aliases: []string{"c"},
				Usage:   "complete a task on the list",
				Action: func(c *cli.Context) error {
					fmt.Println("completed task: ", c.Args().First())
					return nil
				},
			},
			{
				Name:    "template",
				Aliases: []string{"t"},
				Usage:   "options for task templates",
				Subcommands: []cli.Command{
					{
						Name:  "add",
						Usage: "add a new template",
						Action: func(c *cli.Context) error {
							fmt.Println("new task template: ", c.Args().First())
							return nil
						},
					},
					{
						Name:  "remove",
						Usage: "remove an existing template",
						Action: func(c *cli.Context) error {
							fmt.Println("removed task template: ", c.Args().First())
							return nil
						},
					},
				},
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
