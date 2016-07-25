package main

import (
	"net/http"
	"router"
)

func main() {
	http.HandleFunc("/text", router.TextFormatter)
	http.ListenAndServe("localhost:10101", nil)
}
