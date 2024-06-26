package main

import (
	"github.com/dawwson/nomadcoin/blockchain"
)

func main() {
	chain := blockchain.GetBlockChain()
	chain.AddBlock("Second Block");
	chain.AddBlock("Third Block");
	chain.AddBlock("Fourth Block");
	for _, block := range(chain.GetAllBlocks()) {
		block.PrintBlock()
	}
}