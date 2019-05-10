package network

import "github.com/urfave/cli"

var (
	TestnetName = cli.StringFlag{
		Name:  "testnet",
		Usage: "Testnet ",
	}
	Port = cli.IntFlag{
		Name:  "Port",
		Usage: "Port served by all nodes in the testnet",
		Value: 9000,
	}
)
