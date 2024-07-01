package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

const (
	port string = ":3000"
	baseURL string = "http://localhost"
)

type URL string

// NOTE: https://pkg.go.dev/encoding#TextMarshaler
// TextMarshaler 인터페이스를 implements 키워드 없이 암묵적으로 구현합니다.
// URL이 어떻게 JSON으로 마샬링 될 지를 구현합니다.
func (path URL) MarshalText() ([]byte, error) {
	url := fmt.Sprintf("%s%s%s", baseURL, port, path)
	return []byte(url), nil
}

type URLDescription struct {
	// NOTE: struct field tag - struct가 json으로 변환될 때 tag에 지정한 이름으로 변경됨
	URL         URL `json:"url"`
	Method      string `json:"method"`
	Description string `json:"description"`
	Payload     string `json:"payload,omitempty"` // omitempty: 값이 비어있으면 필드 제거
}

func documentation(w http.ResponseWriter, r *http.Request) {
	data := []URLDescription{
		{
			URL: URL("/"),
			Method: "GET",
			Description: "See Documentation",
		},
		{
			URL: URL("/blocks"),
			Method: "POST",
			Description: "Add a Block",
			Payload: "data:string",
		},
	}
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}

func main() {
	http.HandleFunc("/", documentation)
	fmt.Printf("Listening on http://localhost%s", port)
	log.Fatal(http.ListenAndServe(port, nil))
}