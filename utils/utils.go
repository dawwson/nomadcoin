package utils

import (
	"bytes"
	"crypto/sha256"
	"encoding/gob"
	"fmt"
	"log"
)

// 에러 핸들링
func HandleErr(err error) {
	if err != nil {
		log.Panic(err)
	}
}

// interface -> []byte
func ToBytes(i interface{}) []byte {
	var buffer bytes.Buffer
	err := gob.NewEncoder(&buffer).Encode(i)
	HandleErr(err)
	return buffer.Bytes()
}

// []byte -> interface
func FromBytes(i interface{}, data []byte) {
	err := gob.NewDecoder(bytes.NewReader(data)).Decode(i)
	HandleErr(err)
}

// 해시 생성
func Hash(i interface{}) string {
	// 1. 문자열로 변환
	s := fmt.Sprintf("%v", i)
	// 2. 해시 생성
	hash := sha256.Sum256([]byte(s))
	// 3. 16진수로 변환
	return fmt.Sprintf("%x", hash)
}
