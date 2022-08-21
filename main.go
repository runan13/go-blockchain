package main

import (
	"awesomeProject/blockchain"
	"fmt"
)

func main() {

	chain := blockchain.GetBlockchain()
	chain.AddBlock("Second Block")
	chain.AddBlock("Third Block")

	for _, block := range chain.AllBlocks() {
		fmt.Println("-------------------------------------------------------------------")
		fmt.Printf("Data : %s\n", block.Data)
		fmt.Printf("Hash : %s\n", block.Hash)
		fmt.Printf("Prev Hash : %s\n", block.PrevHash)
		fmt.Println("-------------------------------------------------------------------")
	}
}
