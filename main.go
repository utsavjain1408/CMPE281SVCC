package main


import (

	"net/http"

	"io"

)


func main() {

	http.HandleFunc("/", foo)

	http.ListenAndServe(":8080", nil)


	}


	const AddForm = `

	<form method="GET" action="/add">

	<ul><li><b>First</b></li><li>Second</li><li>Third</li></ul>
	
</form>

	`
	

	
	func foo(w http.ResponseWriter, req *http.Request){

	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	io.WriteString(w, AddForm)


	}