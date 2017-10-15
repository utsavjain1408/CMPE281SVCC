	package main

	import (
		"net/http"
		"io"
	)

	func main() {
	http.HandleFunc("/", foo)
	http.ListenAndServe(":8080", nil)

	}



	func foo(w http.ResponseWriter, req *http.Request) {
		io.WriteString(w, "Hello World")
	}
