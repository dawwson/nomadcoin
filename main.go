package main

import "github.com/dawwson/nomadcoin/blockchain"

/*
func (blockchain Blockchain) getLastHash() string {
	if len(blockchain.blocks) > 0 {
		return blockchain.blocks[len(blockchain.blocks) - 1].hash
	}
	return ""
}

func (blockchain *Blockchain) addBlock(data string) {
	// 1. 추가할 블록 생성
	newBlock := Block{data, "", blockchain.getLastHash()}
	// 2. 추가할 블록의 해시 생성(data + 이전 블록의 해시)
	hash := sha256.Sum256([]byte(newBlock.data + newBlock.prevHash))
	// 3. 추가할 블록의 해시 업데이트(16진수 string으로 변환)
	newBlock.hash = fmt.Sprintf("%x", hash)
	// 4. 블록체인에 블록 추가
	blockchain.blocks = append(blockchain.blocks, newBlock)
}

func (blockchain Blockchain) listBlocks() {
	for _, block := range blockchain.blocks {
		fmt.Printf("Data: %s\n", block.data)
		fmt.Printf("Hash: %s\n", block.hash)
		fmt.Printf("Previous Hash: %s\n", block.prevHash)
	}
}
*/
func main() {
	chain := blockchain.GetBlockChain()
}