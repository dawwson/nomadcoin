package rest

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/dawwson/nomadcoin/blockchain"
	"github.com/dawwson/nomadcoin/utils"
	"github.com/gorilla/mux"
)

const baseURL string = "http://localhost"

var port string

type url string

type urlDescription struct {
	// NOTE: struct field tag - struct가 json으로 변환될 때 tag에 지정한 이름으로 변경됨
	URL         url    `json:"url"`
	Method      string `json:"method"`
	Description string `json:"description"`
	Payload     string `json:"payload,omitempty"` // omitempty: 값이 비어있으면 필드 제거
}

type addBlockBody struct {
	Message string
}

type errorResponse struct {
	ErrorMessage string `json:"errorMessage"`
}

// url json 마샬링 지정
func (path url) MarshalText() ([]byte, error) {
	// NOTE: https://pkg.go.dev/encoding#TextMarshaler
	// TextMarshaler 인터페이스를 implements 키워드 없이 암묵적으로 구현
	totalUrl := fmt.Sprintf("%s%s%s", baseURL, port, path)
	return []byte(totalUrl), nil
}

// GET /
func documentation(w http.ResponseWriter, r *http.Request) {
	data := []urlDescription{
		{
			URL:         url("/"),
			Method:      "GET",
			Description: "See Documentation",
		},
		{
			URL:         url("/status"),
			Method:      "GET",
			Description: "See the Status of the Blockchain",
		},
		{
			URL:         url("/blocks"),
			Method:      "GET",
			Description: "See All Blocks",
			Payload:     "data:string",
		},
		{
			URL:         url("/blocks"),
			Method:      "POST",
			Description: "Add a Block",
			Payload:     "data:string",
		},
		{
			URL:         url("/blocks/{hash}"),
			Method:      "GET",
			Description: "See a Block",
		},
	}
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}

// GET /status
func status(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(blockchain.BlockChain())
}

// GET /blocks, POST /blocks
func blocks(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		json.NewEncoder(w).Encode(blockchain.BlockChain().Blocks())
	case http.MethodPost:
		var b addBlockBody
		utils.HandleErr(json.NewDecoder(r.Body).Decode(&b))
		blockchain.BlockChain().AddBlock(b.Message)
		w.WriteHeader(http.StatusCreated)
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

// GET /blocks/{hash}
func block(w http.ResponseWriter, r *http.Request) {
	hash := mux.Vars(r)["hash"]

	block, err := blockchain.FindBlock(hash)
	if err == blockchain.ErrNotFound {
		json.NewEncoder(w).Encode(errorResponse{fmt.Sprint(err)})
	} else {
		json.NewEncoder(w).Encode(block)
	}
}

// Content-Type: application/json 헤더 지정해주는 미들웨어
func jsonContentTypeMiddleware(next http.Handler) http.Handler {
	// NOTE: apater pattern
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")
		next.ServeHTTP(w, r)
	})
}

// ========= Export =========

func Start(at int) {
	port = fmt.Sprintf(":%d", at)

	router := mux.NewRouter()
	router.Use(jsonContentTypeMiddleware)

	router.HandleFunc("/", documentation).Methods("GET")
	router.HandleFunc("/status", status).Methods("GET")
	router.HandleFunc("/blocks", blocks).Methods("GET", "POST")
	router.HandleFunc("/blocks/{hash:[a-f0-9]+}", block).Methods("GET")

	fmt.Printf("Rest Server is listening on http://localhost%s\n", port)
	log.Fatal(http.ListenAndServe(port, router))
}
