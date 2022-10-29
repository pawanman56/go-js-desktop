package main

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"runtime"

	"github.com/zserge/lorca"
)

func getUserName() string {
	var envKey string

	if runtime.GOOS == "windows" {
		envKey = "USERNAME"
	} else {
		envKey = "USER"
	}

	return os.Getenv(envKey)
}

func main() {
	// Pass HTML from data URI
	ui, err := lorca.New("", "", 600, 400)

	ui.Bind("getUserName", getUserName)

	if err != nil {
		log.Fatal(err)
	}

	ln, err := net.Listen("tcp", "127.0.0.1:0")

	if err != nil {
		log.Fatal(err)
	}

	defer ln.Close()

	go http.Serve(ln, http.FileServer(http.Dir("./www")))
	ui.Load(fmt.Sprintf("http://%s", ln.Addr()))

	defer ui.Close()
	<-ui.Done()
}
