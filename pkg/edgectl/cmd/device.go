package cmd

import (
	"fmt"
	"os"

	"github.com/edgexfoundry/edgex-go/pkg/clients/metadata"
	"github.com/edgexfoundry/edgex-go/pkg/clients/types"
	"github.com/spf13/cobra"
)

var DeviceClient metadata.DeviceClient = nil

func InitDeviceCommand() *cobra.Command {
	// Initialize REST client for device
	url := BaseMetadataURL + DeviceUriPath
	params := types.EndpointParams{
		ServiceKey:  MetadataServiceKey,
		Path:        DeviceUriPath,
		UseRegistry: false,
		Url:         url}

	DeviceClient = metadata.NewDeviceClient(params, types.Endpoint{})

	rootCommand := &cobra.Command{
		Use:   "device [COMMANDS]",
		Short: "Operate with device",
	}

	GetCommand := &cobra.Command{
		Use:   "get [OPTIONS]",
		Short: "Get specific device object",
		Run:   GetDeviceByName,
	}

	ListCommand := &cobra.Command{
		Use:   "list",
		Short: "List all device objects",
		Run:   GetAllDevices,
	}

	rootCommand.AddCommand(
		GetCommand,
		ListCommand)

	return rootCommand
}

func AddNewDevice(c *cobra.Command, args []string) {
	// TODO(apopovych): implement proper method in client
}

func UpdateDevice(c *cobra.Command, args []string) {
	// TODO(apopovych): implement proper method in client
}

func GetAllDevices(c *cobra.Command, args []string) {
	devices, err := DeviceClient.Devices()
	if err != nil {
		fmt.Println(err)
	} else {
		w.Init(os.Stdout, 0, 8, 1, '\t', 0)
		fmt.Fprintln(w, "Name\tAdminState\tOperatingState\tLables\t")
		for _, device := range devices {
			fmt.Fprintf(w, "%s\t%s\t%s\t%s\t\n",
				device.Name,
				device.AdminState,
				device.OperatingState,
				device.Labels)
		}
		fmt.Fprintln(w)
		w.Flush()
	}
}

func GetDevicesWithLabel(c *cobra.Command, args []string) {
	// TODO(apopovych): implement proper method in client
}

func GetDeviceByProfileId(c *cobra.Command, args []string) {
	// TODO(apopovych): implement proper method in client
}

func GetDeviceByServiceId(c *cobra.Command, args []string) {
	// TODO(apopovych): implement proper method in client
}

func GetDeviceByServiceName(c *cobra.Command, args []string) {
	// TODO(apopovych): implement proper method in client
}

func GetDeviceByAddressableName(c *cobra.Command, args []string) {
	// TODO(apopovych): implement proper method in client
}

func GetDeviceByProfileName(c *cobra.Command, args []string) {
	// TODO(apopovych): implement proper method in client
}

func GetDeviceByAddressableId(c *cobra.Command, args []string) {
	// TODO(apopovych): implement proper method in client
}

func GetDeviceById(c *cobra.Command, args []string) {
	device, err := DeviceClient.Device(args[0])
	if err != nil {
		fmt.Println(err)
	} else {
		w.Init(os.Stdout, 0, 8, 1, '\t', 0)
		fmt.Fprintln(w, "Name\tAdminState\tOperatingState\t")
		fmt.Fprintf(w, "%s\t%s\t%s\t\n",
			device.Name,
			device.AdminState,
			device.OperatingState)
		fmt.Fprintln(w)
		w.Flush()
	}
}

func DeleteDeviceById(c *cobra.Command, args []string) {
	// TODO(apopovych): implement proper method in client
}

func SetDeviceOpStateById(c *cobra.Command, args []string) {
	// TODO(apopovych): implement proper method in client
}

func SetDeviceAdminStateById(c *cobra.Command, args []string) {
	// TODO(apopovych): implement proper method in client
}

func SetDeviceLastReportedById(c *cobra.Command, args []string) {
	// TODO(apopovych): implement proper method in client
}

func SetDeviceLastReportedByIdNotify(c *cobra.Command, args []string) {
	// TODO(apopovych): implement proper method in client
}

func SetDeviceLastConnectedById(c *cobra.Command, args []string) {
	// TODO(apopovych): implement proper method in client
}

func SetLastConnectedByIdNotify(c *cobra.Command, args []string) {
	// TODO(apopovych): implement proper method in client
}

func CheckForDevice(c *cobra.Command, args []string) {
	// TODO(apopovych): implement proper method in client
}

func GetDeviceByName(c *cobra.Command, args []string) {
	device, err := DeviceClient.DeviceForName(args[0])
	if err != nil {
		fmt.Println(err)
	} else {
		w.Init(os.Stdout, 0, 8, 1, '\t', 0)
		fmt.Fprintln(w, "Name\tAdminState\tOperatingState\tLastConnected\tLastReported\tLabels\t")
		fmt.Fprintf(w, "%s\t%s\t%s\t%d\t%d\t%s\t\n",
			device.Name,
			device.AdminState,
			device.OperatingState,
			device.LastConnected,
			device.LastReported,
			device.Labels)
		fmt.Fprintln(w)
		w.Flush()
	}
}
