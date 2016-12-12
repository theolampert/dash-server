package main

import (
	"./server"
	"github.com/fatih/color"
	"github.com/urfave/cli"
	"os"
	"os/exec"
	"path/filepath"
)

func clear() {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func main() {
	var port string

	var pem string
	var key string

	app := cli.NewApp()
	app.Name = "Dash Server"
	app.Version = "0.0.1"
	app.Usage = "Serve MPEG-DASH from the command-line"
	app.Description = "Small, command-line HTTP/2 file server for serving MPEG-DASH content."

	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:        "port",
			Value:       ":3000",
			Usage:       "Port to be used used for server",
			Destination: &port,
		},

		cli.StringFlag{
			Name:        "pem",
			Usage:       ".pem file for TLS",
			Destination: &pem,
		},

		cli.StringFlag{
			Name:        "key",
			Usage:       ".key file for TLS",
			Destination: &key,
		},
	}

	app.Action = func(c *cli.Context) error {
		workDir, _ := filepath.Abs(c.Args().Get(0))

		if c.NArg() > 0 {
			clear()
			options := make(map[string]string)

			options["workDir"] = workDir
			options["port"] = port

			if pem != "" && key != "" {
				options["pem"] = pem
				options["key"] = key
				clear()
				color.Yellow("Dash-Server started at: https://localhost%s/", options["port"])
			} else {
				clear()
				color.Yellow("Dash-Server started at: http://localhost%s/", options["port"])
			}

			server.Run(options)
		} else {
			color.Red("Please specify a directory to serve.")
		}

		return nil
	}

	app.Run(os.Args)
}
