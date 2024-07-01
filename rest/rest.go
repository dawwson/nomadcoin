package rest

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/dawwson/nomadcoin/blockchain"
	"github.com/dawwson/nomadcoin/utils"
)

const baseURL string = "http://localhost"
var port string

type url string

// NOTE: https://pkg.go.dev/encoding#TextMarshaler
// TextMarshaler 인터페이스를 implements 키워드 없이 암묵적으로 구현합니다.
// URL이 어떻게 JSON으로 마샬링 될 지를 구현합니다.
func (path url) MarshalText() ([]byte, error) {
	totalUrl := fmt.Sprintf("%s%s%s", baseURL, port, path)
	return []byte(totalUrl), nil
}

type urlDescription struct {
	// NOTE: struct field tag - struct가 json으로 변환될 때 tag에 지정한 이름으로 변경됨
	URL         url `json:"url"`
	Method      string `json:"method"`
	Description string `json:"description"`
	Payload     string `json:"payload,omitempty"` // omitempty: 값이 비어있으면 필드 제거
}

type addBlockBody struct {
	Message string
}

func documentation(w http.ResponseWriter, r *http.Request) {
	data := []urlDescription{
		{
			URL: url("/"),
			Method: "GET",
			Description: "See Documentation",
		},
		{
			URL: url("/blocks"),
			Method: "GET",
			Description: "See All Blocks",
			Payload: "data:string",
		},
		{
			URL: url("/blocks"),
			Method: "POST",
			Description: "Add a Block",
			Payload: "data:string",
		},
		{
			URL: url("/blocks/{id}"),
			Method: "GET",
			Description: "See a Block",
		},
	}
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}

func blocks(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
		case http.MethodGet: 
			w.Header().Add("Content-Type", "application/json")
			json.NewEncoder(w).Encode(blockchain.GetBlockChain().GetAllBlocks())
		case http.MethodPost:
			var b addBlockBody 
			utils.HandleErr(json.NewDecoder(r.Body).Decode(&b))
			blockchain.GetBlockChain().AddBlock(b.Message)
			w.WriteHeader(http.StatusCreated)
	}
}

func Start(at int) {
	handler := http.NewServeMux()
	handler.HandleFunc("/", documentation)
	handler.HandleFunc("/blocks", blocks)
	
	port = fmt.Sprintf(":%d", at)
	fmt.Printf("Rest Server is listening on http://localhost%s\n", port)
	log.Fatal(http.ListenAndServe(port, handler))
}