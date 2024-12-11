package shutter

import "github.com/urfave/cli/v2"

type Config struct {
	Enabled bool
}

func ConfigFromCli(ctx *cli.Context) Config {
	return Config{
		Enabled: ctx.Bool(EnabledFlag.Name),
	}
}

var (
	CliFlags = []cli.Flag{
		&EnabledFlag,
	}

	EnabledFlag = cli.BoolFlag{
		Name:  "shutter",
		Usage: "Enable the Shutter private transaction provider (defaults to false)",
	}
)
