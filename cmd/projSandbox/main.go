package main

import (
	"log"
	"net/http"

	"github.com/acom21/pet-proj-dsandbox/service"
)

func main() {
	handler := http.HandlerFunc(service.DefaultServer)
	log.Fatal(http.ListenAndServe(":5000", handler))
}
