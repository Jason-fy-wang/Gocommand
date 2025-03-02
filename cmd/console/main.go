package main

import (
	"os"

	"com.learn/command/trending"
	"com.learn/command/tui"
	log "github.com/sirupsen/logrus"
	cli "github.com/urfave/cli/v2"
)

func main() {
	app := cli.NewApp()

	app.Commands = []*cli.Command{
		tui.TuiCommand,
		trending.GhQuery,
	}

	err := app.Run(os.Args)

	if err != nil {
		log.Infof("app start failed. %v", err)
	}
}
