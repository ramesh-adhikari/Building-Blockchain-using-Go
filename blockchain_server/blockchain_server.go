package main

import (
	"blockchain/block"
	"blockchain/wallet"
	"io"
	"log"
	"net/http"
	"strconv"
)

var cache map[string]*block.Blockchain = make(map[string]*block.Blockchain)

type BlockchainServer struct {
	port uint16
}

func NewBlockchainServer(port uint16) *BlockchainServer {
	return &BlockchainServer{port}
}

func (s *BlockchainServer) Port() uint16 {
	return s.port
}

func (s *BlockchainServer) GetBlockchain() *block.Blockchain {
	bc, ok := cache["blockchain"]
	if !ok {
		minerWallet := wallet.NewWallet()
		bc = block.NewBlockchain(minerWallet.BlockchainAddress(), s.Port())
		cache["blockchain"] = bc
		log.Printf("private_key %v", minerWallet.PrivateKeyStr())
		log.Printf("public_key %v", minerWallet.PublicKeyStr())
		log.Printf("address %v", minerWallet.BlockchainAddress())
	}
	return bc
}

func (s *BlockchainServer) GetChain(w http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case http.MethodGet:
		w.Header().Add("Content-Type", "application/json")
		bc := s.GetBlockchain()
		m, _ := bc.MarshalJSON()
		io.WriteString(w, string(m[:]))
	}
	log.Printf("ERROR: Invalid HTTP Method")
}

func (s *BlockchainServer) Run() {
	http.HandleFunc("/", s.GetChain)
	log.Fatal(http.ListenAndServe("0.0.0.0:"+strconv.Itoa(int(s.Port())), nil))
}
