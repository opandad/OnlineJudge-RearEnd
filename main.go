<<<<<<< HEAD
// main
=======
>>>>>>> 452e7828ec81fcc6920c650941bf76d7e6b0f36b
package main

import (
	"fmt"
<<<<<<< HEAD
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
=======
)

func main() {
	fmt.Print("helloworld")
>>>>>>> 452e7828ec81fcc6920c650941bf76d7e6b0f36b
}
