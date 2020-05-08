package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/axetroy/gmock/internal/app"
)

func main() {

	port := flag.Int("port", 8080, "port of server")
	host := flag.String("host", "localhost", "address of server")

	//demon := flag.Bool("demon", false, "demon mod")

	flag.Parse()

	targetDir := ""

	if len(flag.Args()) > 0 {
		targetDir = flag.Arg(0)
	} else {
		if cwd, err := os.Getwd(); err != nil {
			log.Fatal(err)
		} else {
			targetDir = cwd
		}
	}

	addr := fmt.Sprintf("%s:%d", *host, *port)

	if err := app.Server(addr, targetDir); err != nil {
		log.Fatal(err)
	}
}
