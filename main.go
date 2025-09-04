package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
)

func addHandler(w http.ResponseWriter, r *http.Request) {
	aStr := r.URL.Query().Get("a")
	bStr := r.URL.Query().Get("b")

	a, err1 := strconv.Atoi(aStr)
	b, err2 := strconv.Atoi(bStr)

	if err1 != nil || err2 != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	fmt.Fprintf(w, "Sum = %d\n", a+b)
}

func main() {
	http.HandleFunc("/add", addHandler)
	log.Println("Starting server on :8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
