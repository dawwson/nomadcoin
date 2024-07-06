package blockchain

import (
	"errors"
	"strings"
	"time"

	"github.com/dawwson/nomadcoin/db"
	"github.com/dawwson/nomadcoin/utils"
)

type Block struct {
	Data       string `json:"data"`
	Hash       string `json:"hash"`
	PrevHash   string `json:"prevHash,omitempty"`
	Height     int    `json:"height"`
	Difficulty int    `json:"difficulty"`
	Nonce      int    `json:"nonce"`
	Timestamp  int    `json:"timestamp"`
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
	target := strings.Repeat("0", b.Difficulty)

	for {
		// block 생성 시간 추가
		b.Timestamp = int(time.Now().Unix())
		// block 해시 생성
		hash := utils.Hash(b)

		// 생성한 해시가 target으로 시작하면 block의 hash로 지정
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
		Difficulty: BlockChain().difficulty(),
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
