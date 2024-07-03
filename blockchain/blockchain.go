package blockchain

import (
	"crypto/sha256"
	"errors"
	"fmt"
	"sync"
)

type Block struct {
	Data     string `json:"data"`
	Hash     string `json:"hash"`
	PrevHash string `json:"prevHash,omitempty"`
	Height   int    `json:"height"`
}

type blockchain struct {
	// block pointer slice
	// NOTE: 블록을 추가할 때 전체블록이 복사될 수 있으므로, 포인터로 저장하여 메모리 사용량을 줄인다!
	blocks []*Block 
}

var bc *blockchain
var once sync.Once

func (b *Block ) calculateHash() {
	// 1. 추가할 블록의 해시 생성(data + 이전 블록의 해시)
	hash := sha256.Sum256([]byte(b.Data + b.PrevHash))
	// 2. 추가할 블록의 해시 업데이트(16진수 string으로 변환)
	b.Hash = fmt.Sprintf("%x", hash)
}

func getLastHash() string {
	totalBlocks := GetBlockChain().blocks
	
	if len(totalBlocks) == 0 {
		return ""
	}
	return totalBlocks[len(totalBlocks) - 1].Hash
}

func createBlock(data string) *Block {
	newBlock := Block{data, "", getLastHash(), len(GetBlockChain().blocks) + 1}
	newBlock.calculateHash()
	return &newBlock
}

// ========= Export =========

func (bc *blockchain) AddBlock(data string) {
	bc.blocks = append(bc.blocks, createBlock(data))
}

func (bc *blockchain) AllBlocks() []*Block {
	return bc.blocks
}

var ErrNotFound = errors.New("block not found")

func (bc *blockchain) GetBlock(height int) (*Block, error) {
	if height > len(bc.blocks) {
		return nil, ErrNotFound
	}
	return bc.blocks[height - 1], nil
}

// NOTE: singleton pattern
func GetBlockChain() *blockchain {
	if bc == nil {
		once.Do(func ()  {
			// 블록체인 인스턴스를 한 번만 생성해서 그 메모리 주소를 저장함
			bc = &blockchain{}
			bc.AddBlock("Genesis Block")
		})
	}
	return bc
}