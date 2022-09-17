// Bzipper reads input, bzip2-compresses it, and writes it out.
// Bzipper 读取输入，bzip2 压缩它，然后写出来。
package main

import (
	"gopl2022.io/ch13/bzip"
	"io"
	"log"
	"os"
)

func main() {
	w := bzip.NewWriter(os.Stdout)
	if _, err := io.Copy(w, os.Stdin); err != nil {
		log.Fatalf("bzipper: %v\n", err)
	}
	if err := w.Close(); err != nil {
		log.Fatalf("bzipper: close: %v\n", err)
	}
}

//!-
