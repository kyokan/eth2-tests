package genesis

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

type Node struct {
	Id        string
	LocalId   int
	Ip        string
	TestnetId string
}

func GetNodes(testnetId string) []Node {
	resp, err := http.Get("http://localhost:8000/testnets/" + testnetId + "/nodes")
	if err != nil {
		log.Fatal("Error requesting nodes information", err)
	}
	if resp.StatusCode != 200 {
		log.Fatal("There was an error communicating with the testnet ", err)
	}
	nodesInfo, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal("There was an error reading the response from genesis", err)
	}
	var nodes []Node
	err = json.Unmarshal(nodesInfo, &nodes)
	if err != nil {
		log.Fatal("There was an error reading the response from genesis", err)
	}
	return nodes
}
