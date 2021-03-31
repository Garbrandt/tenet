package main

import (
	"github.com/Garbrandt/tenet/pkg/servers"
)

var (
	Version  = "v1.0.0"
	Revision = "unset"
)

func main() {

	servers.Run()
}
