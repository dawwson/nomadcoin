package utils

import (
	"bytes"
	"encoding/gob"
	"log"
)

func HandleErr(err error) {
	if err != nil {
		log.Panic(err)
	}
}

func ToBytes(i interface{}) []byte {
	var buffer bytes.Buffer
	err := gob.NewEncoder(&buffer).Encode(i)
	HandleErr(err)
	return buffer.Bytes()
}
