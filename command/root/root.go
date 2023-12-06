package root

import (
	"fmt"
	"os"

	"github.com/sigmachain/sigmachain/command/backup"
	"github.com/sigmachain/sigmachain/command/genesis"
	"github.com/sigmachain/sigmachain/command/helper"
	"github.com/sigmachain/sigmachain/command/ibft"
	"github.com/sigmachain/sigmachain/command/license"
	"github.com/sigmachain/sigmachain/command/monitor"
	"github.com/sigmachain/sigmachain/command/peers"
	"github.com/sigmachain/sigmachain/command/secrets"
	"github.com/sigmachain/sigmachain/command/server"
	"github.com/sigmachain/sigmachain/command/status"
	"github.com/sigmachain/sigmachain/command/txpool"
	"github.com/sigmachain/sigmachain/command/version"
	"github.com/sigmachain/sigmachain/command/whitelist"
	"github.com/spf13/cobra"
)

type RootCommand struct {
	baseCmd *cobra.Command
}

func NewRootCommand() *RootCommand {
	rootCommand := &RootCommand{
		baseCmd: &cobra.Command{
			Short: "SIGMA is an Ethereum-compatible Blockchain network",
		},
	}

	helper.RegisterJSONOutputFlag(rootCommand.baseCmd)

	rootCommand.registerSubCommands()

	return rootCommand
}

func (rc *RootCommand) registerSubCommands() {
	rc.baseCmd.AddCommand(
		version.GetCommand(),
		txpool.GetCommand(),
		status.GetCommand(),
		secrets.GetCommand(),
		peers.GetCommand(),
		monitor.GetCommand(),
		ibft.GetCommand(),
		backup.GetCommand(),
		genesis.GetCommand(),
		server.GetCommand(),
		whitelist.GetCommand(),
		license.GetCommand(),
	)
}

func (rc *RootCommand) Execute() {
	if err := rc.baseCmd.Execute(); err != nil {
		_, _ = fmt.Fprintln(os.Stderr, err)

		os.Exit(1)
	}
}
