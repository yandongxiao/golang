package main

import (
	"fmt"
	stdlog "log"
	"net/http"
	"time"

	"github.com/yandongxiao/golang-learning/lib/context/log"
)

func main() {
	http.HandleFunc("/", log.Decorate(handler))
	stdlog.Fatal(http.ListenAndServe("127.0.0.1:8080", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log.Println(ctx, "handler started")
	defer log.Println(ctx, "handler ended")
	select {
	case <-time.After(5 * time.Second):
		_, _ = fmt.Fprintln(w, "hello")
	case <-ctx.Done():
		err := ctx.Err()
		log.Println(ctx, err.Error())
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
}
