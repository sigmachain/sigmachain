package ibft

import (
	"github.com/sigmachain/sigmachain/command/helper"
	"github.com/sigmachain/sigmachain/command/ibft/candidates"
	"github.com/sigmachain/sigmachain/command/ibft/propose"
	"github.com/sigmachain/sigmachain/command/ibft/quorum"
	"github.com/sigmachain/sigmachain/command/ibft/snapshot"
	"github.com/sigmachain/sigmachain/command/ibft/status"
	_switch "github.com/sigmachain/sigmachain/command/ibft/switch"
	"github.com/spf13/cobra"
)

func GetCommand() *cobra.Command {
	ibftCmd := &cobra.Command{
		Use:   "ibft",
		Short: "Top level IBFT command for interacting with the IBFT consensus. Only accepts subcommands.",
	}

	helper.RegisterGRPCAddressFlag(ibftCmd)

	registerSubcommands(ibftCmd)

	return ibftCmd
}

func registerSubcommands(baseCmd *cobra.Command) {
	baseCmd.AddCommand(
		// ibft status
		status.GetCommand(),
		// ibft snapshot
		snapshot.GetCommand(),
		// ibft propose
		propose.GetCommand(),
		// ibft candidates
		candidates.GetCommand(),
		// ibft switch
		_switch.GetCommand(),
		// ibft quorum
		quorum.GetCommand(),
	)
}
