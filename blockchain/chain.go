package blockchain

import (
	"sync"

	"github.com/dawwson/nomadcoin/db"
	"github.com/dawwson/nomadcoin/utils"
)

type blockchain struct {
	LatestHash string `json:"latestHash"`
	Height     int    `json:"height"`
}

var bc *blockchain
var once sync.Once

func (bc *blockchain) persist() {
	db.SaveBlockchain(utils.ToBytes(bc))
}

// ========= Export =========

func (bc *blockchain) AddBlock(data string) {
	block := createBlock(data, bc.LatestHash, bc.Height+1)
	bc.LatestHash = block.Hash
	bc.Height = block.Height
	bc.persist()
}

// NOTE: singleton pattern
func BlockChain() *blockchain {
	if bc == nil {
		once.Do(func() {
			// 블록체인 인스턴스를 한 번만 생성해서 그 메모리 주소를 저장함
			bc = &blockchain{"", 0}
			bc.AddBlock("Genesis Block")
		})
	}
	return bc
}
