package genesis

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

type TestnetParameter struct {
}

type TestnetResource struct {
	Cpus   string
	Memory string
}

type Testnet struct {
	Servers    []int
	Blockchain string
	Nodes      int
	Images     []string
	Resources  []TestnetResource
	Params     TestnetParameter
}

func DeployTestnet(blockchain string, images []string, nodes int, output string) {
	testNet := Testnet{
		[]int{1},
		blockchain,
		nodes,
		images,
		[]TestnetResource{
			{"", ""},
		},
		TestnetParameter{},
	}
	json, err := json.Marshal(testNet)
	if err != nil {
		log.Fatal("Error preparing testnet configuration", err)
	}
	log.Printf(string(json))
	resp, err := http.Post("http://localhost:8000/testnets", "application/json", bytes.NewBuffer(json))
	if err != nil {
		log.Fatal("Error sending a testnet configuration", err)
	}
	if resp.StatusCode != 200 {
		log.Fatal("There was an error deploying the testnet", err)
	}
	testnetId, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal("There was an error reading the response from genesis", err)
	}
	log.Printf("Testnet deployed with id %s", testnetId)
	if output != "" {
		err := ioutil.WriteFile(output, testnetId, 0644)
		if err != nil {
			log.Fatal("There was an error saving testnet id to file", err)
		}
	}
}
