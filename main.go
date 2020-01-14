package main

import (
	"os"

	"github.com/lp74/uc74-blockchain-go/cli"
)

func main() {
	defer os.Exit(0)
	cmd := cli.CommandLine{}
	cmd.Run()
}
