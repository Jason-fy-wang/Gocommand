package trending

import (
	cli "github.com/urfave/cli/v2"
)

var GhQuery = &cli.Command{
	Name:  "gh",
	Usage: "gihug trending",
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:     "lang",
			Usage:    "specific language to query. f.g: java, python,go,javascript",
			Value:    "go",
			Required: false,
		},
		&cli.StringFlag{
			Name:  "format",
			Usage: "output format , f.g: json, table",
			Value: "table",
		},
	},
	Action: func(context *cli.Context) error {
		inputLang := context.String("lang")
		format := context.String("format")
		GhTrendingQuery(inputLang, format)
		return nil
	},
}
