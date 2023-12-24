package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
		d, err := io.ReadAll(r.Body)

		if err != nil {
			http.Error(rw, "oops", http.StatusBadRequest)
		}

		fmt.Printf("Data: %s\n", d)
	})

	http.ListenAndServe(":3000", nil)
}
