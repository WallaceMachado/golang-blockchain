package main

import (
	"github.com/WallaceMachado/golang-blockchain/cli"
	"os"
)

func main() {
	defer os.Exit(0)
	cmd := cli.CommandLine{}
	cmd.Run()
}
