package blockchain

import (
	"crypto/sha256"
	"fmt"

	"github.com/dawwson/nomadcoin/db"
	"github.com/dawwson/nomadcoin/utils"
)

type Block struct {
	Data     string `json:"data"`
	Hash     string `json:"hash"`
	PrevHash string `json:"prevHash,omitempty"`
	Height   int    `json:"height"`
}

func (b *Block) persist() {
	db.SaveBlock(b.Hash, utils.ToBytes(b))
}

func createBlock(data string, prevHash string, height int) *Block {
	// 블록 및 해시 생성
	block := Block{
		Data:     data,
		Hash:     "",
		PrevHash: prevHash,
		Height:   height,
	}
	payload := block.Data + block.PrevHash + fmt.Sprint(block.Height)
	block.Hash = fmt.Sprintf("%x", sha256.Sum256(([]byte(payload))))
	// DB에 블록 저장
	block.persist()
	return &block
}
