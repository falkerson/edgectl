package cmd

import (
	"text/tabwriter"
)

const (
	BaseMetadataURL    = "http://edgex.webzoom.top:48081"
	MetadataServiceKey = "edgex-core-metadata"

	AddressableUriPath = "/api/v1/addressable"
	DeviceUriPath      = "/api/v1/device"
)

var w = new(tabwriter.Writer)
