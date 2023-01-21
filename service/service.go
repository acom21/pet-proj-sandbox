package service

import (
	"fmt"
	"net/http"
)

func DefaultServer(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, 20)
}
