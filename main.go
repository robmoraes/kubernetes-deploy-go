package main

import (
	"fmt"
	"log"
	"net/http"
)

const webContent = "dev-ops-ninja:v100"

func main() {
	http.HandleFunc("/", helloHandler)
	log.Fatal(http.ListenAndServe(":80", nil))
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, webContent)
}
