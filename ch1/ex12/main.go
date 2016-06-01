package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
)

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	if cyclesString := r.URL.Query().Get("cycles"); cyclesString != "" {
		cycles, err := strconv.Atoi(cyclesString)
		if err != nil {
			fmt.Fprintf(os.Stdout, "Server: %v\n", err)
		}
		lissajous(w, cycles)
	} else {
		lissajous(w, 5)
	}
}
