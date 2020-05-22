package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/axetroy/gmock/lib/daemon"
)

const defaultPort = 8080
const defaultHost = "0.0.0.0"

func getPort() int {
	PORT := os.Getenv("PORT")

	if PORT == "" {
		return defaultPort
	}

	if port, err := strconv.Atoi(PORT); err != nil {
		return defaultPort
	} else {
		return port
	}
}

func getHost() string {
	HOST := os.Getenv("HOST")

	if HOST == "" {
		return defaultHost
	}

	return HOST
}

func getDir() string {
	var targetDir string = os.Getenv("DIR")

	if targetDir != "" {
		return targetDir
	}

	if len(flag.Args()) > 0 {
		targetDir = flag.Arg(0)
	} else {
		if cwd, err := os.Getwd(); err != nil {
			log.Fatal(err)
		} else {
			targetDir = cwd
		}
	}

	return targetDir
}

func main() {
	port := flag.Int("port", getPort(), "port of server")
	host := flag.String("host", getHost(), "address of server")
	isDaemonMode := flag.Bool("daemon", false, "enable daemon mod")

	flag.Parse()

	targetDir := getDir()

	addr := fmt.Sprintf("%s:%d", *host, *port)

	err := daemon.Start(func() error {
		if err := Server(addr, targetDir); err != nil {
			return err
		}

		return nil
	}, *isDaemonMode)

	if err != nil {
		log.Fatalln(err)
	}
}
