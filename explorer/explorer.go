package explorer

import (
	"fmt"
	"log"
	"net/http"
	"text/template"

	"github.com/dawwson/nomadcoin/blockchain"
)

const templateDir string = "explorer/templates/"

var templates *template.Template

type homeData struct {
	PageTitle string
	Blocks    []*blockchain.Block
}

// GET / - home 템플릿 로딩
func home(w http.ResponseWriter, r *http.Request) {
	data := homeData{"Home", nil}
	templates.ExecuteTemplate(w, "home", data)
}

// GET /add - add 템플릿 로딩 &
// FIXME: POST /add - 블록 추가 (REST API랑 중복?)
func add(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		templates.ExecuteTemplate(w, "add", nil)
	case "POST":
		r.ParseForm()
		data := r.Form.Get("blockData")
		blockchain.BlockChain().AddBlock(data)
		http.Redirect(w, r, "/", http.StatusSeeOther)
	}
}

// ========= Export =========

func Start(port int) {
	// load templates
	// NOTE: Must - 에러가 발생하면 출력해줌
	templates = template.Must(template.ParseGlob(templateDir + "pages/*.gohtml"))
	templates = template.Must(templates.ParseGlob(templateDir + "partials/*.gohtml"))

	handler := http.NewServeMux()
	handler.HandleFunc("/", home)
	handler.HandleFunc("/add", add)

	fmt.Printf("Explorer is listening on http://localhost:%d\n", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), handler))
}
