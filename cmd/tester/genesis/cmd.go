package genesis

import (
	"io/ioutil"
	"log"
	"net/http"

	"github.com/kyokan/eth2-tests/tester/genesis"
	"github.com/urfave/cli"
)

var (
	GenesisCommand = cli.Command{
		Name:        "genesis",
		Usage:       "Commands to set up genesis",
		Description: `Commands to set up and test genesis`,
		Subcommands: []cli.Command{
			cli.Command{
				Name:        "up",
				Usage:       "Tests genesis is available",
				Description: `Tests that genesis is available on port 8000`,
				Action:      testGenesisAvailable,
				Flags:       []cli.Flag{},
			},
			cli.Command{
				Name:        "testnet",
				Usage:       "Deploys a new testnet",
				Description: `Deploys a new testnet to genesis`,
				Action:      deployTestnet,
				Flags: []cli.Flag{
					FileOutputFlag,
					BlockchainFlag,
					NumberOfNodesFlag,
				},
			},
		},
	}
)

func deployTestnet(ctx *cli.Context) {
	blockchain := ctx.String(BlockchainFlag.Name)
	output := ctx.String(FileOutputFlag.Name)
	genesis.DeployTestnet(blockchain, genesis.Images[blockchain], ctx.Int(NumberOfNodesFlag.Name), output)
}

func testGenesisAvailable(ctx *cli.Context) {
	resp, err := http.Get("http://localhost:8000/servers")
	if err != nil {
		log.Fatal("There was an error contacting genesis", err)
	}
	defer resp.Body.Close()
	_, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal("There was an error reading the response from genesis", err)
	}
	if resp.StatusCode != 200 {
		log.Fatalf("The genesis server returned an error: %d", resp.StatusCode)
	}
	log.Printf("Genesis server contacted successfully")
}
