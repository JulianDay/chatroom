package server

import (
	"log"
	"net/http"
)

func Start(addr string) {
	RegisterHandle()
	log.Fatal(http.ListenAndServe(addr, nil))
}
