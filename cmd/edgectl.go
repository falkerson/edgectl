package main

import (
	"github.com/falkerson/edgectl/pkg/edgectl"
)

func main() {
	command := edgectl.NewRootCommand()
	edgectl.Execute(command)
}
