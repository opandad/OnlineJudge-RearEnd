// main
package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strings"
)

func PrintServerInfo(r *http.Request) {
	r.ParseForm()
	fmt.Println(r.Form)
	fmt.Println("path: ", r.URL.Path)
	fmt.Println("scheme: ", r.URL.Scheme)
	fmt.Println(r.Form["url_long"])
	for k, v := range r.Form {
		fmt.Println("key: ", k)
		fmt.Println("val: ", strings.Join(v, ""))
	}
}

//主页
func Index(w http.ResponseWriter, r *http.Request) {
	PrintServerInfo(r)
	t := template.New("some template")

	t, _ = t.ParseFiles("dist/index.html", nil)
	t.Execute(w)
}

func main() {
	http.HandleFunc("/dist", Index)
	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
