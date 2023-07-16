package main

import (
	"blockchain/wallet"
	"fmt"
	"log"
)

func init() {
	log.SetPrefix("Blockchain: ")
}

func main() {
	w := wallet.NewWallet()
	fmt.Println(w.PrivateKeyStr())
	fmt.Println(w.PublicKeyStr())
	fmt.Println(w.BlockchainAddress())

	t := wallet.NewTransaction(w.PrivateKey(), w.PublicKey(), w.BlockchainAddress(), "B", 1.0)
	fmt.Printf("signature %s \n", t.GenerateSignature())
	// myBlockchainAddress := "my_blockchain_address"
	// blockChain := NewBlockchain(myBlockchainAddress)
	// blockChain.Print()

	// blockChain.AddTransaction("A", "B", 10)
	// blockChain.AddTransaction("A", "B", 10)
	// blockChain.Mining()
	// previousHash := blockChain.LastBlock().Hash()
	// nonce := blockChain.ProofOfWork()
	// blockChain.CreateBlock(nonce, previousHash)
	// blockChain.Print()

	// fmt.Printf("B %.1f\n", blockChain.CalculateTotalAmount("B"))
	// fmt.Printf("A %.1f\n", blockChain.CalculateTotalAmount("A"))
	// previousHash1 := blockChain.LastBlock().Hash()
	// blockChain.CreateBlock(6, previousHash1)
	// blockChain.Print()

	// previousHash2 := blockChain.LastBlock().Hash()
	// blockChain.CreateBlock(7, previousHash2)
	// blockChain.Print()
	// blockChain.CreateBlock(1, "hash 1")
	// blockChain.Print()
	// block := &Block{nonce: 1}
	// fmt.Printf("%x\n", block.Hash())
}
