package main

import (
	"log"
	"net/http"
)

func main() {
	fs := http.FileServer(http.Dir("./shared"))
	http.Handle("/", fs)

	log.Print("[-] sglfs Listening on :3000...")
	err := http.ListenAndServe(":3000", nil)
	if err != nil {
		log.Fatal(err)
	}
}
