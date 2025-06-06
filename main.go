package main

import (
	"fmt"
	"os"

	"github.com/btschwartz12/xplr/cmds"
)

func main() {
	if err := cmds.New().Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
