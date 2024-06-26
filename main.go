package main

import (
	"fmt"
	"log"
	"net/http"
	"text/template"

	"github.com/dawwson/nomadcoin/blockchain"
)

const port string = ":3000"

type homeData struct {
	PageTitle string
	Blocks    []*blockchain.Block
}

// NOTE: w - response, r - request(포인터, 실제 reqeust 사용)
// NOTE: Must - 에러가 발생하면 출력해줌
func home(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("templates/pages/home.gohtml"))
	data := homeData{"Home", blockchain.GetBlockChain().GetAllBlocks()}
	tmpl.Execute(w, data)
}

func main() {
	// GET /
	http.HandleFunc("/", home)
	fmt.Printf("Listening on http://localhost%s\n", port)
	log.Fatal(http.ListenAndServe(port, nil))
}