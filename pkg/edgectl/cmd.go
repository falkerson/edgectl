package edgectl

import (
	"fmt"
	"os"

	"github.com/falkerson/edgectl/pkg/edgectl/cmd"
	"github.com/spf13/cobra"
)

func NewRootCommand() *cobra.Command {
	rootCmd := &cobra.Command{
		Use:   "edgectl [COMMANDS]",
		Short: "edgectl CLI based client for edgex platform",
	}

	rootCmd.AddCommand(
		addVersion(),
		cmd.InitAddressableCommand(),
		cmd.InitDeviceCommand())

	return rootCmd
}

func addVersion() *cobra.Command {
	return &cobra.Command{
		Use:   "version",
		Short: "Print the version number of edgectl",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Edgectl version v0.1")
		},
	}
}

func Execute(command *cobra.Command) {
	if err := command.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
