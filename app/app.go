package app

import "github.com/urfave/cli"

// Generate will return the line command app ready to be executed.
func Generate() *cli.App {
	app := cli.NewApp()
	app.Name = "Command Line Application"
	app.Usage = "Search IPs and Server names on internet"

	return app
}
