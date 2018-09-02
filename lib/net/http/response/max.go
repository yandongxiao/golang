package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/", func(rw http.ResponseWriter, req *http.Request) {
		rc := http.MaxBytesReader(rw, req.Body, 2)
		defer rc.Close()

		// request body too large
		fmt.Println(io.Copy(os.Stdout, rc))
	})

	http.ListenAndServe(":8080", nil)
}
