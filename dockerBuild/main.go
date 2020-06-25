package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", hello)

	http.ListenAndServe(":80", nil)
}

func hello(w http.ResponseWriter, req *http.Request) {

	req.Header.Set("Content-Type", "text/html")
	fmt.Fprintf(w,
		`<html>
		<body>
		hello
		</body>
		</html>`)
}
