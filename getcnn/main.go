package main

import (
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/jcoppis/goCNNTop10"
)

func main() {
	http.HandleFunc("/", home)
	http.HandleFunc("/about", about)
	http.HandleFunc("/cnn", getCNN)
	fmt.Printf("Servidor iniciado en el puerto '%d'\n", 8080)
	http.ListenAndServe(":8080", nil)
}

func home(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "Mi pagina")
}

func getCNN(w http.ResponseWriter, req *http.Request) {
	items, err := goCNNTop10.Get()
	if err != nil {
		log.Fatal(err)
	}

	for _, item := range items {
		//fmt.Println(item)
		fmt.Fprintln(w, item)
	}
}

func about(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "about")
}
