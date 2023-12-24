package handlers

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

type Data struct {
	l *log.Logger
}

func NewData(l *log.Logger) *Data {
	return &Data{l}
}

func (h *Data) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	h.l.Println("Testing")

	d, err := io.ReadAll(r.Body)

	if err != nil {
		http.Error(rw, "oops", http.StatusBadRequest)
	}

	fmt.Printf("Data: %s\n", d)
}
