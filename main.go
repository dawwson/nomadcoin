package main

import (
	"github.com/dawwson/nomadcoin/explorer"
	"github.com/dawwson/nomadcoin/rest"
)

func main() {
	go explorer.Start(3000)
	rest.Start(4000)
}