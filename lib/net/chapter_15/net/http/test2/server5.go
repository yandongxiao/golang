package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func uploadHandler(writer http.ResponseWriter, request *http.Request) {
	file, header, err := request.FormFile("file")

	if err != nil {
		fmt.Fprintln(writer, err)

		return
	}

	defer file.Close()

	out, err := os.Create("/tmp/uploadedfile")

	if err != nil {
		fmt.Fprintln(writer, err)

		return
	}

	defer out.Close()

	_, err = io.Copy(out, file)

	if err != nil {
		fmt.Fprintln(writer, err)
	}

	fmt.Fprintf(writer, "File uploaded successfully : ")
	fmt.Fprintf(writer, header.Filename)
}

func main() {
	http.HandleFunc("/", uploadHandler)

	http.ListenAndServe(":8080", nil)
}
