package main

import (
	"github.com/dawwson/nomadcoin/cli"
	"github.com/dawwson/nomadcoin/db"
)

func main() {
	defer db.Close()
	cli.Start()
}
