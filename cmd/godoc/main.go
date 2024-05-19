package main

import (
	"fmt"
	"github.com/misuaaki/godoc"
	gen "github.com/misuaaki/godoc/app"
	"github.com/misuaaki/godoc/app/formatter"
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
			Name:  "test-scan",
			Usage: "Scan your project at the given path and under and populate the Packages map - used for testing purposes",
			Action: func(c *cli.Context) error {
				path := c.Args().First()
				if path == "" {
					return cli.ShowCommandHelp(c, "test-scan")
				}

				extension := c.Args().Get(1)
				if extension == "" {
					return cli.ShowCommandHelp(c, "test-scan")
				}

				s := scanner.NewScanner()
				if err := s.Scan(path, extension); err != nil {
					return err
				}

				fmt.Println(s.Packages[path].String())
				return nil
			},
		},
		{
			Name:  "test-generate",
			Usage: "Generate a simple textual of your documentation - testing purposes only",
			Action: func(c *cli.Context) error {
				path := c.Args().First()
				if path == "" {
					return cli.ShowCommandHelp(c, "test-generate")
				}

				err := gen.GenerateTextualDocumentation(path)
				if err != nil {
					return err
				}
				return nil
			},
		},
		{
			Name:  "init",
			Usage: "Generate a formatted PDF of your documentation",
			Action: func(c *cli.Context) error {
				path := c.Args().First()
				if path == "" {
					return cli.ShowCommandHelp(c, "generate")
				}

				f := formatter.NewFormatter()
				if err := f.GeneratePDF(path); err != nil {
					return err
				}
				return nil
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
