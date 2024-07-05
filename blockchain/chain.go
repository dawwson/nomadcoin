package blockchain

import (
	"bytes"
	"encoding/gob"
	"fmt"
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

func (bc *blockchain) restore(data []byte) {
	err := gob.NewDecoder(bytes.NewReader(data)).Decode(bc)
	utils.HandleErr(err)
}

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

// NOTE: singleton pattern - 블록체인 인스턴스를 한 번만 생성
func BlockChain() *blockchain {
	if bc == nil {
		once.Do(func() {
			bc = &blockchain{"", 0}
			fmt.Printf("LastestHash: %s\nHeight: %d\n", bc.LatestHash, bc.Height)
			//	db에서 checkpoint 조회
			checkpoint := db.Checkpoint()

			if checkpoint == nil {
				// checkpoint가 없으면 genesis 블록으로 블록체인 초기화
				bc.AddBlock("Genesis Block")
			} else {
				// checkpoint가 있으면 복원
				fmt.Println("\n🚀 Restoring...")
				bc.restore(checkpoint)
			}
		})
	}
	fmt.Printf("LastestHash: %s\nHeight: %d\n", bc.LatestHash, bc.Height)
	return bc
}
