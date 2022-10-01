package main

import (
	"awesomeProject/blockchain"
	"awesomeProject/cli"
	"awesomeProject/db"
)

func main() {
	blockchain.Blockchain()
	defer db.Close()
	cli.Start()
}
