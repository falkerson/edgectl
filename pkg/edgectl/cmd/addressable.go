package cmd

import (
	"fmt"
	"os"
	"text/tabwriter"

	"github.com/edgexfoundry/edgex-go/pkg/clients/metadata"
	"github.com/edgexfoundry/edgex-go/pkg/clients/types"
	"github.com/spf13/cobra"
)

const (
	BaseMetadataURL    = "http://edgex.webzoom.top:48081"
	AddressableUriPath = "/api/v1/addressable"
	MetadataServiceKey = "edgex-core-metadata"
)

var (
	client metadata.AddressableClient = nil
	w                                 = new(tabwriter.Writer)
)

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
