package main

import (
	"log"
	"os"

	"github.com/kyokan/eth2-tests/cmd/tester/genesis"
	"github.com/kyokan/eth2-tests/cmd/tester/network"
	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()
	app.Version = "1.0.0"
	app.Name = "test"
	app.Usage = "Tests a testnet of clients"
	app.Commands = []cli.Command{
		genesis.GenesisCommand,
		network.NetworkCommand,
	}
	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
