package main

import (
	"github.com/dawwson/nomadcoin/cli"
	"github.com/dawwson/nomadcoin/db"
)

func main() {
	defer db.Close()
	cli.Start()

	// difficulty := 5
	// target := strings.Repeat("0", difficulty)
	// nonce := 1

	// for {
	// 	hash := fmt.Sprintf("%x", sha256.Sum256([]byte("Hello"+fmt.Sprint(nonce))))

	// 	fmt.Printf("Hash: %s\nTarget: %s\nNonce: %d\n\n", hash, target, nonce)

	// 	if strings.HasPrefix(hash, target) {
	// 		return
	// 	} else {
	// 		// 조건이 맞지 않으면 nonce를 늘려서 반복
	// 		nonce++
	// 	}
	// }
}
