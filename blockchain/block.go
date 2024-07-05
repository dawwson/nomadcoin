package blockchain

import (
	"crypto/sha256"
	"errors"
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

// 블록 저장
func (b *Block) persist() {
	db.SaveBlock(b.Hash, utils.ToBytes(b))
}

// 블록 복원
func (b *Block) restore(data []byte) {
	utils.FromBytes(b, data)
}

// 블록 생성
func createBlock(data string, prevHash string, height int) *Block {
	// 1. 블록 객체 생성
	block := Block{
		Data:     data,
		Hash:     "",
		PrevHash: prevHash,
		Height:   height,
	}

	// 2. 해시 생성
	payload := block.Data + block.PrevHash + fmt.Sprint(block.Height)
	block.Hash = fmt.Sprintf("%x", sha256.Sum256(([]byte(payload))))

	// 3. DB에 블록 저장
	block.persist()

	return &block
}

// ========= Export =========

var ErrNotFound = errors.New("block not found")

// 해시로 블록 조회
func FindBlock(hash string) (*Block, error) {
	blockBytes := db.Block(hash)

	if blockBytes == nil {
		return nil, ErrNotFound
	}

	block := &Block{}
	block.restore(blockBytes)

	return block, nil
}
