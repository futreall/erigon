package shutter

import (
	"embed"
	"encoding/json"
	"fmt"

	"github.com/urfave/cli/v2"
)

type Config struct {
	Enabled                          bool     `json:"-"`
	InstanceId                       uint64   `json:"instanceId"`
	SequencerContractAddress         string   `json:"sequencerContractAddress"`
	ValidatorRegistryContractAddress string   `json:"validatorRegistryContractAddress"`
	KeyBroadcastContractAddress      string   `json:"keyBroadcastContractAddress"`
	KeyperSetManagerContractAddress  string   `json:"keyperSetManagerContractAddress"`
	KeyperBootnodes                  []string `json:"keyperBootnodes"`
}

func ConfigFromCli(ctx *cli.Context) Config {
	enabled := ctx.Bool(EnabledFlag.Name)
	if !enabled {
		return Config{}
	}

	chainName := ctx.String("chain") // TODO - reuse global chain name flag (needs refactoring due to circ dep)
	config := readConfig(chainName)
	config.Enabled = enabled

	// check for cli overrides
	if ctx.IsSet(KeyperBootnodesFlag.Name) {
		config.KeyperBootnodes = ctx.StringSlice(KeyperBootnodesFlag.Name)
	}

	return config
}

//go:embed chainconfigs
var chainConfigs embed.FS

func readConfig(chainName string) Config {
	fileName := fmt.Sprintf("chainconfigs/%s.json", chainName)
	f, err := chainConfigs.Open(fileName)
	if err != nil {
		panic(fmt.Sprintf("could not open shutter chain config for %s: %v", fileName, err))
	}
	defer func() {
		_ = f.Close()
	}()

	config := Config{}
	decoder := json.NewDecoder(f)
	err = decoder.Decode(&config)
	if err != nil {
		panic(fmt.Sprintf("could not parse shutter chain config for %s: %v", fileName, err))
	}

	return config
}

var (
	CliFlags = []cli.Flag{
		&EnabledFlag,
		&KeyperBootnodesFlag,
	}

	EnabledFlag = cli.BoolFlag{
		Name:  "shutter",
		Usage: "Enable the Shutter encrypted transactions provider (defaults to false)",
	}

	KeyperBootnodesFlag = cli.StringSliceFlag{
		Name:  "shutter.keyper.bootnodes",
		Usage: "Use to override the default keyper bootnodes (defaults to using the bootnodes from the embedded config)",
	}
)
