package explorer

import (
	"fmt"
	"log"
	"net/http"
	"text/template"

	"github.com/dawwson/nomadcoin/blockchain"
)

const (
	port        string = ":3000"
	templateDir string = "explorer/templates/"
)

var templates *template.Template

type homeData struct {
	PageTitle string
	Blocks    []*blockchain.Block
}

// NOTE: w - response, r - request(포인터, 실제 reqeust 사용)
func home(w http.ResponseWriter, r *http.Request) {
	data := homeData{"Home", blockchain.GetBlockChain().GetAllBlocks()}
	templates.ExecuteTemplate(w, "home", data)
}

func add(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		templates.ExecuteTemplate(w, "add", nil)
	case "POST": 
		r.ParseForm()
		data := r.Form.Get("blockData")
		blockchain.GetBlockChain().AddBlock(data)
		http.Redirect(w, r, "/", http.StatusSeeOther)
	}
}

func Start() {
		// load templates
	// NOTE: Must - 에러가 발생하면 출력해줌
	templates = template.Must(template.ParseGlob(templateDir + "pages/*.gohtml"))
	templates = template.Must(templates.ParseGlob(templateDir + "partials/*.gohtml"))

	http.HandleFunc("/", home)
	http.HandleFunc("/add", add)
	fmt.Printf("Listening on http://localhost%s\n", port)
	log.Fatal(http.ListenAndServe(port, nil))
}