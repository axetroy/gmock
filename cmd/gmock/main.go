package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/axetroy/gmock/internal/app"
	"github.com/axetroy/gmock/internal/lib/daemon"
)

func main() {

	port := flag.Int("port", 8080, "port of server")
	host := flag.String("host", "localhost", "address of server")
	isDaemonMode := flag.Bool("daemon", false, "enable daemon mod")

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

	err := daemon.Start(func() error {
		if err := app.Server(addr, targetDir); err != nil {
			return err
		}

		return nil
	}, *isDaemonMode)

	if err != nil {
		log.Fatalln(err)
	}
}
