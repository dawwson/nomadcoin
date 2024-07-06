package blockchain

import (
	"fmt"
	"sync"

	"github.com/dawwson/nomadcoin/db"
	"github.com/dawwson/nomadcoin/utils"
)

const (
	defaultDifficulty  int = 2 // 기본 난이도
	difficultyInterval int = 5 // 난이도 계산하는 블록 간격
	blockInterval      int = 2 // 블록 1개 생성 시간
	allowedRange       int = 2 // 블록 생성시간 허용 범위(-2 ~ +2)
)

var bc *blockchain
var once sync.Once

type blockchain struct {
	LatestHash        string `json:"latestHash"`
	Height            int    `json:"height"`
	CurrentDifficulty int    `json:"currentDifficulty"`
}

// 블록체인 저장
func (bc *blockchain) persist() {
	db.SaveCheckpoint(utils.ToBytes(bc))
}

// 블록체인 복원
func (bc *blockchain) restore(data []byte) {
	utils.FromBytes(bc, data)
}

// 채굴 난이도 다시 계산
func (bc *blockchain) recalculateDifficulty() int {
	allBlocks := bc.Blocks()
	// 가장 최근에 생성된 블록
	lastestBlock := allBlocks[0]
	// 가장 마지막에 난이도가 계산된 블록
	lastRecalculatedBlock := allBlocks[difficultyInterval-1]
	// 두 블록의 실제 생성 시간 간격(분 단위)
	actualBlockInterval := (lastestBlock.Timestamp / 60) - (lastRecalculatedBlock.Timestamp / 60)
	// 예상 블록 생성 시간 간격
	expectedBlockInterval := difficultyInterval * blockInterval

	// 예상 시간 범위를 넘어서면 난이도를 높이고, 넘지 않으면 난이도를 낮춤.
	// 그 외의 경우 현재 수준 유지
	if actualBlockInterval <= (expectedBlockInterval - allowedRange) {
		return bc.CurrentDifficulty + 1
	} else if actualBlockInterval >= (expectedBlockInterval + allowedRange) {
		return bc.CurrentDifficulty - 1
	} else {
		return bc.CurrentDifficulty
	}
}

// 블록체인 난이도 조회
// : 블록 높이 5의 배수마다 난이도를 다시 계산
func (bc *blockchain) difficulty() int {
	if bc.Height == 0 {
		return defaultDifficulty
	} else if bc.Height%difficultyInterval == 0 {
		return bc.recalculateDifficulty()
	} else {
		return bc.CurrentDifficulty
	}
}

// ========= Export =========

// 블록체인에 블록 추가
func (bc *blockchain) AddBlock(data string) {
	block := createBlock(data, bc.LatestHash, bc.Height+1)
	bc.LatestHash = block.Hash
	bc.Height = block.Height
	bc.CurrentDifficulty = block.Difficulty
	bc.persist()
}

// 블록체인의 모든 블록 조회(최신순)
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

// 블록체인 생성 또는 불러오기
func BlockChain() *blockchain {
	// NOTE: singleton pattern - 블록체인 인스턴스를 한 번만 생성
	if bc == nil {
		once.Do(func() {
			bc = &blockchain{
				Height: 0,
			}
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
