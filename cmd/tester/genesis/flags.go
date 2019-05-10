package genesis

import "github.com/urfave/cli"

var (
	FileOutputFlag = cli.StringFlag{
		Name:  "file",
		Usage: "File to write the testnet ID to ",
		Value: "",
	}
	BlockchainFlag = cli.StringFlag{
		Name:  "blockchain",
		Usage: "Blockchain ",
		Value: "prysm",
	}
	NumberOfNodesFlag = cli.IntFlag{
		Name:  "numNodes",
		Usage: "Number of nodes to deploy",
		Value: 3,
	}
)
