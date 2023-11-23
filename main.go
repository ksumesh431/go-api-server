package main

import (
	"fmt"
	"log"
	"net/http"
)

func formHandler(resp http.ResponseWriter, req *http.Request) {

	if err := req.ParseForm(); err != nil {
		fmt.Fprintf(resp, "ParseForm() err: %x", err)
		return
	}
	fmt.Fprintf(resp, "POST request successful")

	name := req.FormValue("name")
	address := req.FormValue("address")
	fmt.Fprintf(resp, "\nName = %s\n", name)
	fmt.Fprintf(resp, "Address = %s\n", address)

}

func helloHandler(resp http.ResponseWriter, req *http.Request) {
	if req.URL.Path != "/hello" {
		http.Error(resp, "404 not found", http.StatusNotFound)
		return
	}
	if req.Method != "GET" {
		http.Error(resp, "method is not supported", http.StatusNotFound)
		return
	}
	fmt.Fprintf(resp, "Hola!")
}

func main() {
	// http.fileserver looks for index.html in the path given in fileserver
	fileserver := http.FileServer(http.Dir("./static"))
	http.Handle("/", fileserver)

	http.HandleFunc("/formPost", formHandler)

	http.HandleFunc("/form", func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "/form.html", http.StatusSeeOther)
	})

	http.HandleFunc("/hello", helloHandler)

	fmt.Printf("Starting server at port 8080\n")

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
