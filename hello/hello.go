package main

import (
	"os"
	"io"
	"log"
	"net/http"
)

func HelloServer(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "Ahoj, Hexo!\n")
}

func main() {
	http.HandleFunc("/", HelloServer)
	err := http.ListenAndServe(os.Args[1], nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
