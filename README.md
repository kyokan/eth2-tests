# Ethereum 2.0 tests

This repository contains a series of sanity tests to be built against Ethereum 2.0 clients.

## `tester`

The tester program connects to an existing Genesis server and deploys, maintains and tests a testnet.

See https://www.github.com/Whiteblock/genesis to set up a testnet server.

### Check genesis is available locally

`tester genesis up`

### Create a testnet

`tester genesis tesnet --blockchain [prysm|]`

Optionally, you can indicate a file to store the testnet ID.


`tester genesis tesnet --blockchain [prysm|] -f out`

### Check all nodes in the testnet can serve traffic on a port

`tester network --testnet <testnetId>`

# Chains supported

| Name | Image                        | Repository                              |
|------|------------------------------|-----------------------------------------|
|Prysm |gcr.io/whiteblock/prysm:master| https://github.com/prysmaticlabs/prysm  | 


