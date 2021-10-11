package http_client_server

import (
	"io"
	"net/http"
)

func hello(res http.ResponseWriter, req *http.Request) {
	res.Header().Set(
		"Content-Type",
		"text/html",
	)
	io.WriteString(
		res,
		`<DOCTYPE html>
	<html>
	<head>
	<title>Hello, World</title>
	</head>
	<body>
	Hello, World!
	</body>
	</html>`,
	)
}

type ItWorks struct{}

func (p *ItWorks) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "It Works!\n")
}

func Http_client_server_test() {
	http.HandleFunc("/hello", hello)
	//iw := new(ItWorks)
	//http.ListenAndServe(":9000", iw)
	http.ListenAndServe(":9000", nil)
}
