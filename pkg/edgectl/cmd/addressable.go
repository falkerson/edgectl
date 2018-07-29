package cmd

import (
	"fmt"
	"os"

	"github.com/edgexfoundry/edgex-go/pkg/clients/metadata"
	"github.com/edgexfoundry/edgex-go/pkg/clients/types"
	"github.com/spf13/cobra"
)

var AddressableClient metadata.AddressableClient = nil

func InitAddressableCommand() *cobra.Command {
	// Initialize REST client for addressable
	url := BaseMetadataURL + AddressableUriPath
	params := types.EndpointParams{
		ServiceKey:  MetadataServiceKey,
		Path:        AddressableUriPath,
		UseRegistry: false,
		Url:         url}

	AddressableClient = metadata.NewAddressableClient(
		params, types.Endpoint{})

	rootCommand := &cobra.Command{
		Use:   "addressable [COMMANDS]",
		Short: "Operate with addressable",
	}

	GetCommand := &cobra.Command{
		Use:   "get [OPTIONS]",
		Short: "Get specific addressable object",
		Run:   GetAddressableByName,
	}

	rootCommand.AddCommand(GetCommand)

	return rootCommand
}

func GetAddressableByName(c *cobra.Command, args []string) {
	addressable, err := AddressableClient.AddressableForName(args[0])
	if err != nil {
		fmt.Println(err)
	} else {
		w.Init(os.Stdout, 0, 8, 1, '\t', 0)
		fmt.Fprintln(w, "Name\tProtocol\tHTTPMethod\tAddress\tPort\t")
		fmt.Fprintf(w, "%s\t%s\t%s\t%s\t%d\t\n",
			addressable.Name,
			addressable.Protocol,
			addressable.HTTPMethod,
			addressable.Address,
			addressable.Port)
		fmt.Fprintln(w)
		w.Flush()
	}
}

func GetAllAddressables(c *cobra.Command, args []string) {
	// TODO(apopovych): implement proper method in client
}

func AddAddressable(c *cobra.Command, args []string) {
	// TODO(apopovych): implement proper method in client
}

func UpdateAddressable(c *cobra.Command, args []string) {
	// TODO(apopovych): implement proper method in client
}

func GetAddressableById(c *cobra.Command, args []string) {
	// TODO(apopovych): implement proper method in client
}

func DeleteAddressableById(c *cobra.Command, args []string) {
	// TODO(apopovych): implement proper method in client
}

func DeleteAddressableByName(c *cobra.Command, args []string) {
	// TODO(apopovych): implement proper method in client
}

func GetAddressableByTopic(c *cobra.Command, args []string) {
	// TODO(apopovych): implement proper method in client
}

func GetAddressableByPort(c *cobra.Command, args []string) {
	// TODO(apopovych): implement proper method in client
}

func GetAddressableByPublisher(c *cobra.Command, args []string) {
	// TODO(apopovych): implement proper method in client
}

func GetAddressableByAddress(c *cobra.Command, args []string) {
	// TODO(apopovych): implement proper method in client
}
