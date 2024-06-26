package main

import (
	"fmt"
	"log"
	"net/http"
)

const port string = ":3000"

func main() {
	// GET /
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// NOTE: w - response, r - request(포인터, 실제 reqeust 사용)
		fmt.Fprint(w, "Hello from home!")
	})
	fmt.Printf("Listening on http://localhost%s\n", port)
	log.Fatal(http.ListenAndServe(port, nil))
}