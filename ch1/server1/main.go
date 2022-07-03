//Server1 is a minimal "echo" server
// Server1 是一个最小的“回声”服务器。

package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {

	http.HandleFunc("/", handler) // each request calls handler
	log.Fatal(http.ListenAndServe("localhost:8000", nil))

}

//handler echoes the Path component of the request URL r
func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "URL.Path = %q\n", r.URL.Path)
}
