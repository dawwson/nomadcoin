package main

import (
	"github.com/dawwson/nomadcoin/blockchain"
)

func main() {
	blockchain.BlockChain().AddBlock("First")
	blockchain.BlockChain().AddBlock("Second")
	blockchain.BlockChain().AddBlock("Third")
}
