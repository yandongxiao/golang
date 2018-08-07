package main

import (
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
)

//  Multiple programs calling TempDir simultaneously will not choose the same directory.
// 在协程内并发使用可能是有问题的
func main() {
	content := []byte("temporary file's content")
	dir, err := ioutil.TempDir("/tmp", "example")
	if err != nil {
		log.Fatal(err)
	}

	defer os.RemoveAll(dir) // clean up

	tmpfn := filepath.Join(dir, "tmpfile")
	if err := ioutil.WriteFile(tmpfn, content, 0666); err != nil {
		log.Fatal(err)
	}
}
