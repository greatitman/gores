package main

import (
	."fmt"
	"net/http"
	"strings"
	"log"
)

func sayhelloName(w http.ResponseWriter,r *http.Request) {
	r.ParseForm()
	Println(r.Form)
	Println("path",r.URL.Path)
	Println("scheme", r.URL.Scheme)
	Println(r.Form["url_long"])
	
	for k, v := range r.Form {
		Println("key:",k)
		Println("val:",strings.Join(v,""))
	}
	Fprintf(w,"hello http!")
}

func main(){
	http.HandleFunc("/",sayhelloName)
	err := http.ListenAndServe(":9090",nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
