package shutter

import (
	"flag"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/urfave/cli/v2"
)

func TestConfigFromCli(t *testing.T) {
	flagSet := flag.NewFlagSet("test", flag.ContinueOnError)
	flagSet.String("chain", "", "")
	flagSet.Bool("shutter", false, "")
	err := flagSet.Set("chain", "chiado")
	require.NoError(t, err)
	err = flagSet.Set("shutter", "true")
	require.NoError(t, err)
	ctx := cli.NewContext(cli.NewApp(), flagSet, nil)
	chiadoConfig := ConfigFromCli(ctx)
	require.True(t, chiadoConfig.Enabled)
	require.Equal(t, uint64(102000), chiadoConfig.InstanceId)
	require.Equal(t, "0x2aD8E2feB0ED5b2EC8e700edB725f120576994ed", chiadoConfig.SequencerContractAddress)
	require.Equal(t, "0xa9289A3Dd14FEBe10611119bE81E5d35eAaC3084", chiadoConfig.ValidatorRegistryContractAddress)
	require.Equal(t, "0x9D31865BEffcE842FBd36CDA587aDDA8bef804B7", chiadoConfig.KeyBroadcastContractAddress)
	require.Equal(t, "0xC4DE9FAf4ec882b33dA0162CBE628B0D8205D0c0", chiadoConfig.KeyperSetManagerContractAddress)
	wantKeyperBootnodes := []string{
		"/ip4/167.99.177.227/tcp/23005/p2p/12D3KooWSdm5guPBdn8DSaBphVBzUUgPLg9sZLnazEUrcbtLy254",
		"/ip4/159.89.15.119/tcp/23005/p2p/12D3KooWPP6bp2PJQR8rUvG1SD4qNH4WFrKve6DMgWThyKxwNbbH",
	}
	require.Equal(t, wantKeyperBootnodes, chiadoConfig.KeyperBootnodes)
}

func TestConfigFromCliWithKeyperBootnodesOverride(t *testing.T) {
	flagSet := flag.NewFlagSet("test", flag.ContinueOnError)
	flagSet.String("chain", "", "")
	flagSet.Bool("shutter", false, "")
	flagSet.Var(cli.NewStringSlice(), "shutter.keyper.bootnodes", "")
	err := flagSet.Set("chain", "chiado")
	require.NoError(t, err)
	err = flagSet.Set("shutter", "true")
	require.NoError(t, err)
	err = flagSet.Set("shutter.keyper.bootnodes", "xxx,yyy")
	require.NoError(t, err)
	ctx := cli.NewContext(cli.NewApp(), flagSet, nil)
	chiadoConfig := ConfigFromCli(ctx)
	require.Equal(t, []string{"xxx", "yyy"}, chiadoConfig.KeyperBootnodes)
}
