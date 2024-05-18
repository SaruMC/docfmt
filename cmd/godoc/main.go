package main

import (
	"fmt"
	"github.com/misuaaki/godoc"
	"github.com/misuaaki/godoc/app/scanner"
	"github.com/urfave/cli/v2"
	"log"
	"os"
)

func main() {
	app := cli.NewApp()
	app.Name = "app"
	app.Usage = "A CLI tool that generates a formatted PDF of your documentation"
	app.Version = godoc.Version

	app.Commands = []*cli.Command{
		{
			Name:  "scan",
			Usage: "Scan your project at the given path and under and populate the Packages map - used for testing purposes",
			Action: func(c *cli.Context) error {
				path := c.Args().First()
				if path == "" {
					return cli.ShowCommandHelp(c, "scan")
				}

				s := scanner.NewScanner()
				if err := s.Scan(path, ".go"); err != nil {
					return err
				}

				fmt.Println(s.Packages[path].String())
				return nil
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
