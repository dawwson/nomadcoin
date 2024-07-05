package cli

import (
	"flag"
	"fmt"
	"os"
	"runtime"

	"github.com/dawwson/nomadcoin/explorer"
	"github.com/dawwson/nomadcoin/rest"
)

func usage() {
	fmt.Printf("💛 Welcome to Nomad Coin\n\n")
	fmt.Printf("Please use the following flags\n\n")
	fmt.Printf("-port:   Set the PORT of the server\n")
	fmt.Printf("-mode:   Choose between 'html' and 'rest'\n\n")

	// defer 함수 실행 후 프로그램 종료
	runtime.Goexit()
}

func Start() {
	if len(os.Args) == 1 {
		usage()
	}

	port := flag.Int("port", 4000, "Set port of the server")
	mode := flag.String("mode", "rest", "Choose between 'html' and 'rest'")

	flag.Parse()

	switch *mode {
	case "rest":
		rest.Start(*port)
	case "html":
		explorer.Start(*port)
	default:
		usage()
	}
	fmt.Println(*port, *mode)
}
