package main

import (
	"fmt"
	"net/http"
	"log"
    "strings"
    "strconv"
)

var data[]string

func serveFile(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "./static/index.html")
}

func getData() (string) {
	var output string
	for i, s := range data {
		output += fmt.Sprintf("<button hx-post=\"/remove/%d\" hx-target=\"#results\">X</button><h1>%s</h1>", i, s)
	}

	return output
}

func handleRemove(w http.ResponseWriter, r *http.Request) {
    idStr := strings.TrimPrefix(r.URL.Path, "/remove/")
    id, err := strconv.Atoi(idStr)
    if err != nil {
        log.Fatal(err)
    }
    data = append(data[:id], data[id+1:]...)
    output := getData()
    fmt.Fprintf(w, "%s", output)
}


func handleClick(w http.ResponseWriter, r *http.Request) {
	value := r.FormValue("input")
	data = append(data, string(value))
	output := getData()
	fmt.Fprintf(w, "%s", output)
}

func main() {
	http.HandleFunc("/static", serveFile)
	http.HandleFunc("/add", handleClick)
	http.HandleFunc("/remove/", handleRemove)
	fmt.Println("Listening on port 6969")
	err := http.ListenAndServe(":6969", nil)
	if err != nil {
		log.Fatal(err)
	}
}
