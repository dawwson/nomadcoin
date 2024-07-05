package main

import (
	"github.com/dawwson/nomadcoin/blockchain"
	"github.com/dawwson/nomadcoin/cli"
)

func main() {
	blockchain.BlockChain()
	cli.Start()
}
