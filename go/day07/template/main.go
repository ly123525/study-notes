package main

import (
	"fmt"
	"html/template"
	"net/http"
)

type UserInfo struct {
	Name string
	Sex  string
	Age  int
}

func indexPage(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("./index.html")
	if err != nil {
		fmt.Fprintf(w, "load login.html failed")
		return
	}

	user := UserInfo{
		Name: "Mary",
		Sex:  "男",
		Age:  18,
	}

	t.Execute(w, user)
}

func main() {
	http.HandleFunc("/", indexPage)
	err := http.ListenAndServe(":8000", nil)
	if err != nil {
		fmt.Printf("listen server failed, err:%v\n", err)
		return
	}
}
