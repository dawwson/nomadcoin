package blockchain

import (
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

// 블록체인 저장
func (bc *blockchain) persist() {
	db.SaveCheckpoint(utils.ToBytes(bc))
}

// 블록체인 복원
func (bc *blockchain) restore(data []byte) {
	utils.FromBytes(bc, data)
}

// ========= Export =========

// 블록체인에 블록 추가
func (bc *blockchain) AddBlock(data string) {
	block := createBlock(data, bc.LatestHash, bc.Height+1)
	bc.LatestHash = block.Hash
	bc.Height = block.Height
	bc.persist()
}

// 블록체인의 모든 블록 조회
func (bc *blockchain) Blocks() []*Block {
	var blocks []*Block
	hashCursor := bc.LatestHash

	// 이전 해시 값으로 블록을 하나씩 역으로 조회
	for {
		block, _ := FindBlock(hashCursor)
		blocks = append(blocks, block)

		if block.PrevHash != "" {
			hashCursor = block.PrevHash
		} else {
			break
		}
	}

	return blocks
}

// 블록체인 불러오기
func BlockChain() *blockchain {
	// NOTE: singleton pattern - 블록체인 인스턴스를 한 번만 생성
	if bc == nil {
		once.Do(func() {
			bc = &blockchain{"", 0}
			//	db에서 checkpoint 조회
			checkpoint := db.Checkpoint()

			if checkpoint == nil {
				// checkpoint가 없으면 genesis 블록으로 블록체인 초기화
				bc.AddBlock("Genesis Block")
			} else {
				// checkpoint가 있으면 복원
				fmt.Println("🚀 Restoring...")
				bc.restore(checkpoint)
			}
		})
	}
	return bc
}
