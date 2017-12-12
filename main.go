package main

import ("fmt"
"net/http")

func index_handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>This is Main Page</h1>")
	fmt.Println("main page")
}

func about_handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>This is About Page</h1>")
	fmt.Println("about page")
}

func handlerICon(w http.ResponseWriter, r *http.Request) {}

func main() {
	http.HandleFunc("/favicon.ico", handlerICon) 
	http.HandleFunc("/", index_handler)
	http.HandleFunc("/about/", about_handler)
	http.ListenAndServe(":8000", nil) 
}