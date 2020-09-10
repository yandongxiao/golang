package server

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		log.Println("handler started")
		defer log.Println("handler ended")
		ctx := r.Context()
		select {
		case <-time.After(5 * time.Second):
			_, _ = fmt.Fprintln(w, "hello")
		case <-ctx.Done():
			err := ctx.Err()
			log.Println(err)
			http.Error(w, err.Error(), http.StatusBadRequest)
		}
	})
	log.Fatal(http.ListenAndServe("127.0.0.1:8080", nil))
}
