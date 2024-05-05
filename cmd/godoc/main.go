package godoc

import (
	"github.com/misuaaki/godoc"
	"github.com/urfave/cli/v2"
	"log"
	"os"
)

func main() {
	app := cli.NewApp()
	app.Name = godoc.Name
	app.Usage = godoc.Usage
	app.Version = godoc.Version

	app.Commands = []*cli.Command{
		{
			Name: "",
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
