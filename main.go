package main

import (
	"log"
	"net/url"

	"github.com/zserge/lorca"
)

func main() {
	// Pass HTML from data URI
	ui, err := lorca.New("data:text/html,"+url.PathEscape(`
		<html>
			<head>
			<title>Lorca App</title>
			</head>
			<body>
				<h1 style="padding-top: 40vh; text-align: center;">Hello, Lorca!</h1>
			</body>
		</html>
	`), "", 600, 400)

	if err != nil {
		log.Fatal(err)
	}

	defer ui.Close()
	<-ui.Done()
}
