package main

import (
	"crypto/sha256"
	"fmt"
)

type Block struct {
	data string
	hash string
	prevHash string
}

func main() {
	// 1. 새로 추가할 블록을 생성한다.
	block := Block{"data", "", ""}
	// 2. [data + 이전 해시]를 byte 타입의 slice로 변환하여 해싱한다.
	hash := sha256.Sum256([]byte(block.data + block.prevHash))
	// 3. 해시를 16진수 string으로 변환한다.
	hexHash := fmt.Sprintf("%x", hash)
	// 4. 블록의 해시에 저장한다.
	block.hash = hexHash
	fmt.Println(block)
}