package main

import (
	"fmt"
	"log"
	"net"
	"os"

	"github.com/urfave/cli"
)

func main() {
	app := &cli.App{
		Name:  "Simple network tool",
		Usage: "Lets you query Ips, cname, mx records and name servers",
		Commands: []cli.Command{
			{
				Name:  "ns",
				Usage: "Looks up name servers for particular host",
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:  "host",
						Value: "google.com",
					},
				},
				Action: func(c *cli.Context) error {
					ns, err := net.LookupNS(c.String("host"))
					if err != nil {
						return err
					}
					for i := 0; i < len(ns); i++ {
						fmt.Println(ns[i].Host)
					}
					return nil
				},
			},
			{
				Name:  "nscname",
				Usage: "Looks up cname for particular host",
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:  "host",
						Value: "google.com",
					},
				},
				Action: func(c *cli.Context) error {
					cname, err := net.LookupCNAME(c.String("host"))
					if err != nil {
						return err
					}
					fmt.Println(cname)
					return nil
				},
			},
			{
				Name:  "mx",
				Usage: "Looks up mx records for particular host",
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:  "host",
						Value: "google.com",
					},
				},
				Action: func(c *cli.Context) error {
					mx, err := net.LookupMX(c.String("host"))
					if err != nil {
						return err
					}
					for i := 0; i < len(mx); i++ {
						fmt.Println(mx[i].Host, mx[i].Pref)
					}
					return nil
				},
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}

}
