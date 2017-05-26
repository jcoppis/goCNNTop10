package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"

	"github.com/jcoppis/goCNNTop10"
)

func infoServer(puerto int, endPoints []string) {
	fmt.Printf("Servidor iniciado en http://localhost:%d\n", puerto)
	fmt.Print("Puede conectarse a los siguientes endPoints:")
	for _, endPoint := range endPoints {
		fmt.Printf(" /%s", endPoint)
	}
}

func main() {
	http.HandleFunc("/", home)
	http.HandleFunc("/about", about)
	http.HandleFunc("/cnn", getCNN)
	http.HandleFunc("/json", getJSON)
	infoServer(8080, strings.Fields("about cnn json"))
	log.Fatal(http.ListenAndServe(":8080", nil))
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

func getJSON(w http.ResponseWriter, r *http.Request) {
	item := goCNNTop10.Item{Title: "Mi JSON", URL: "http://json.org", Description: "Ejemplo json"}
	//jsonStr := []byte(`{"title":"Buy cheese and bread for breakfast.", "subtitle":"hola"}`)
	json.NewEncoder(w).Encode(item)
}
