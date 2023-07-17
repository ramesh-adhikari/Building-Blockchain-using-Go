package main

import (
	"blockchain/block"
	"blockchain/utility"
	"blockchain/wallet"
	"bytes"
	"encoding/json"
	"html/template"
	"io"
	"log"
	"net/http"
	"path"
	"strconv"
)

const TEMP_DIR = "templates"

type WalletServer struct {
	port uint16
	// gateway to connect other server
	gateway string
}

func NewWalletServer(port uint16, gateway string) *WalletServer {
	return &WalletServer{port, gateway}
}

func (ws *WalletServer) Port() uint16 {
	return uint16(ws.port)
}

func (ws *WalletServer) Gateway() string {
	return ws.gateway
}

func (ws *WalletServer) Index(w http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case http.MethodGet:
		t, _ := template.ParseFiles(path.Join(TEMP_DIR, "index.html"))
		t.Execute(w, "")
	default:
		log.Printf("ERROR: Invalid HTTP Method")
	}
}

func (ws *WalletServer) Wallet(w http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case http.MethodPost:
		w.Header().Add("Content-Type", "application/json")
		myWallet := wallet.NewWallet()
		m, _ := myWallet.MarshalJSON()
		io.WriteString(w, string(m[:]))
	default:
		w.WriteHeader(http.StatusBadRequest)
		log.Println("ERROR: Invalid HTTP Method")

	}
}

func (ws *WalletServer) CreateTransaction(w http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case http.MethodPost:
		decoder := json.NewDecoder(req.Body)
		var t wallet.TransactionRequest
		error := decoder.Decode(&t)
		if error != nil {
			log.Printf("ERROR: %v", error)
			// io.WriteString(w, string(utils.JsonStatus("fail")))
			return
		}
		if !t.Validate() {
			log.Println("ERROR: missing field(s)")
			return
		}
		publicKey := utility.PublicKeyFromString(*t.SenderPublicKey)
		privateKey := utility.PrivateKeyFromString(*t.SenderPrivateKey, publicKey)
		value, err := strconv.ParseFloat(*t.Value, 32)

		if err != nil {
			log.Println("ERROR: amount parse error")
			return
		}

		value32 := float32(value)

		w.Header().Add("Content-Type", "application/json")

		transaction := wallet.NewTransaction(privateKey, publicKey, *t.SenderBlockchainAddress, *t.RecipientBlockchainAddress, value32)
		signature := transaction.GenerateSignature()
		signatureStr := signature.String()

		bt := &block.TransactionRequest{
			t.SenderBlockchainAddress,
			t.RecipientBlockchainAddress,
			t.SenderPublicKey,
			&value32, &signatureStr,
		}

		m, _ := json.Marshal(bt)
		buf := bytes.NewBuffer(m)

		resp, _ := http.Post(ws.Gateway()+"/transactions", "application/json", buf)

		if resp.StatusCode == 201 {
			log.Println("SUCCESS: transaction send from wallet server to blockchain server")
			return
		}
		log.Println("ERROR: Error occurs when sending transaction  from wallet server to blockchain server")

	default:
		w.WriteHeader(http.StatusBadRequest)
		log.Println("ERROR: Invalid HTTP Method")

	}
}

func (ws *WalletServer) Run() {
	http.HandleFunc("/", ws.Index)
	http.HandleFunc("/wallet", ws.Wallet)
	http.HandleFunc("/transaction", ws.CreateTransaction)
	log.Fatal(http.ListenAndServe("0.0.0.0:"+strconv.Itoa(int(ws.Port())), nil))
}
