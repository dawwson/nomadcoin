package blockchain

import (
	"crypto/sha256"
	"fmt"
	"sync"
)

type Block struct {
	data     string
	hash     string
	prevHash string
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
	hash := sha256.Sum256([]byte(b.data + b.prevHash))
	// 2. 추가할 블록의 해시 업데이트(16진수 string으로 변환)
	b.hash = fmt.Sprintf("%x", hash)
}

func getLastHash() string {
	totalBlocks := GetBlockChain().blocks
	
	if len(totalBlocks) == 0 {
		return ""
	}
	return totalBlocks[len(totalBlocks) - 1].hash
}

func createBlock(data string) *Block {
	newBlock := Block{data, "", getLastHash()}
	newBlock.calculateHash()
	return &newBlock
}

// ========= Export =========

func (bc *blockchain) AddBlock(data string) {
	bc.blocks = append(bc.blocks, createBlock(data))
}

func (bc *blockchain) GetAllBlocks() []*Block {
	return bc.blocks
}

func (b *Block ) PrintBlock() {
	fmt.Println("====================")
	fmt.Printf("Data: %s\n", b.data)
	fmt.Printf("Hash: %s\n", b.hash)
	fmt.Printf("Previous Hash: %s\n", b.prevHash)
}

func GetBlockChain() *blockchain {
	if bc == nil {
		once.Do(func ()  {
			// 블록체인 인스턴스를 한 번만 생성해서 그 메모리 주소를 저장함
			bc = &blockchain{}
			bc.AddBlock("First Block")
		})
	}
	return bc
}