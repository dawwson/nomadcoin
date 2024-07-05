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
