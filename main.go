package main

import (
    "fmt"
	"net/http"
	"text/template"
)

func main() {
	http.HandleFunc("/movein", movein)
	http.HandleFunc("/moveout", moveout)
	http.ListenAndServe("0.0.0.0:5556", nil)
}

func movein(w http.ResponseWriter, req *http.Request) {
	mquery := req.URL.Query()
	name := mquery.Get("name")
	fmt.Println(name)
	t, _ := template.ParseFiles("movein.html")
	t.Execute(w, name)
}

func moveout(w http.ResponseWriter, req *http.Request) {
	mquery := req.URL.Query()
	name := mquery.Get("name")
	fmt.Println(name)
	t, _ := template.ParseFiles("moveout.html")
	t.Execute(w, name)
}
