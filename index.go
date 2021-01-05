package main

import (
	"fmt"
	"math"
	"log"
	"net/http"
)

func reallySlowGreeting() string {
	x := 0.0001
	for i := 0; i <=1000000; i++ {
		x += math.Sqrt(x)
	}
	return "Code.education Rocks!"
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, reallySlowGreeting())
}

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}