package cmd

import (
	"fmt"

	"github.com/edgexfoundry/edgex-go/pkg/clients/metadata"
	"github.com/edgexfoundry/edgex-go/pkg/clients/types"
	"github.com/spf13/cobra"
)

const (
	BaseMetadataURL    = "http://edgex.webzoom.top:48081"
	AddressableUriPath = "/api/v1/addressable"
	MetadataServiceKey = "edgex-core-metadata"
)

var client metadata.AddressableClient = nil

func InitAddressableCommand() *cobra.Command {
	// Initialize REST client for addressable
	url := BaseMetadataURL + AddressableUriPath
	params := types.EndpointParams{
		ServiceKey:  MetadataServiceKey,
		Path:        AddressableUriPath,
		UseRegistry: false,
		Url:         url}

	client = metadata.NewAddressableClient(params, types.Endpoint{})

	rootAddressableCommand := &cobra.Command{
		Use:   "addressable [COMMANDS]",
		Short: "Operate with addressable",
	}

	AddressbleGetCommand := &cobra.Command{
		Use:   "get [OPTIONS]",
		Short: "Get specific addressable object",
		Run:   GetAddressableByName,
	}

	rootAddressableCommand.AddCommand(AddressbleGetCommand)

	return rootAddressableCommand
}

func GetAddressableByName(c *cobra.Command, args []string) {
	fmt.Sprintf("Fetch addressable by name %s", args[0])
	addressable, err := client.AddressableForName(args[0])
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(addressable)
	}
}
