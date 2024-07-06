package blockchain

import (
	"crypto/sha256"
	"errors"
	"fmt"
	"strings"

	"github.com/dawwson/nomadcoin/db"
	"github.com/dawwson/nomadcoin/utils"
)

const difficulty int = 2 // 0이 2개로 시작하는 hash를 찾는다.

type Block struct {
	Data       string `json:"data"`
	Hash       string `json:"hash"`
	PrevHash   string `json:"prevHash,omitempty"`
	Height     int    `json:"height"`
	Difficulty int    `json:"difficulty"`
	Nonce      int    `json:"nonce"`
}

// 블록 저장
func (b *Block) persist() {
	db.SaveBlock(b.Hash, utils.ToBytes(b))
}

// 블록 복원
func (b *Block) restore(data []byte) {
	utils.FromBytes(b, data)
}

// 마이닝 : 조건(difficulty 만큼의 0으로 시작)에 맞는 해시를 생성하는 nonce를 찾는다.
func (b *Block) mine() {
	target := strings.Repeat("0", difficulty)

	for {
		// block을 string으로 변환하여 해시 생성
		blockAsString := fmt.Sprint(b)
		hash := fmt.Sprintf("%x", sha256.Sum256([]byte(blockAsString)))
		fmt.Printf("Block as String: %s\nHash: %s\nTarget: %s\nNonce: %d\n\n\n", blockAsString, hash, target, b.Nonce)

		// hash의 prefix가 target이면 그 해시를 블록의 해시로 지정
		if strings.HasPrefix(hash, target) {
			b.Hash = hash
			break
		} else {
			b.Nonce++
		}
	}
}

// 블록 생성
func createBlock(data string, prevHash string, height int) *Block {
	// 1. 블록 객체 생성
	block := Block{
		Data:       data,
		Hash:       "",
		PrevHash:   prevHash,
		Height:     height,
		Difficulty: difficulty,
		Nonce:      0,
	}

	// 2. 해시 생성
	// payload := block.Data + block.PrevHash + fmt.Sprint(block.Height)
	// block.Hash = fmt.Sprintf("%x", sha256.Sum256(([]byte(payload))))
	block.mine()

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
