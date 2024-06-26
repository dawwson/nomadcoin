package blockchain

type block struct {
	data string
	hash string
	prevHash string
}

type blockchain struct {
	blocks []block
}

var bc *blockchain

func GetBlockChain() *blockchain {
	if bc == nil {
		bc = &blockchain{}
	}
	return bc
}