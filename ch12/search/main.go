// Search is a demo of the params.Unpack function.
// 搜索是 params.Unpack 函数的演示。
package main

import (
	"fmt"
	"gopl2022.io/ch12/params"
	"log"
	"net/http"
)

// search implements the /search URL endpoint.
// 搜索实现 /search URL 端点。
func search(resp http.ResponseWriter, req *http.Request) {
	var data struct {
		/*		Labels     []string `http:"l"`
				MaxResults int      `http:"max"`
				Exact      bool     `http:"x"`*/
		Labels     []string `http:"l"`
		MaxResults int      `http:"max"`
		Exact      bool     `http:"x"`
	}
	data.MaxResults = 10 //set default
	if err := params.Unpack(req, &data); err != nil {
		http.Error(resp, err.Error(), http.StatusBadRequest) //400
		return
	}
	// ...rest of handler...
	// ...其余的处理程序...
	fmt.Fprintf(resp, "Search: %+v\n", data)
}

//!-
func main() {
	http.HandleFunc("/search", search)
	log.Fatal(http.ListenAndServe(":12345", nil))
}

/*
//!+output
./fetch 'http://localhost:12345/search?l=golang&l=programming'
$ ./search &
$ ./fetch 'http://localhost:12345/search'
Search: {Labels:[] MaxResults:10 Exact:false}
$ ./fetch 'http://localhost:12345/search?q=hello&x=123'
x: strconv.ParseBool: parsing "123": invalid syntax
$ ./fetch 'http://localhost:12345/search?q=hello&max=lots'
max: strconv.ParseInt: parsing "lots": invalid syntax
//!-output


*/
