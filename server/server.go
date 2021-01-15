package server

import (
	"OnlineJudge-RearEnd/features/users"
	"fmt"
	"log"
	"net/http"
)

func InitServer() {
	http.HandleFunc("/", index)
	err := http.ListenAndServe(":9000", nil)
	if err != nil {
		log.Fatal("Listen and server:", err)
	}
}

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hello go")

	//test
	users.LoginVerifyByEmail("abc@qq.com", "abc")
}
