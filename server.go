package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"
)

func main() {

	filename := "out.txt"
	if len(os.Args) > 1 {
		filename = os.Args[1]
	}

	output := "Default Output"
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Printf("Could not read data from file %s: %s", filename, err)
	} else {
		output = string(data)
	}

	srv := httptest.NewTLSServer((http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, output)
	})))
	defer srv.Close()

	fmt.Printf("Listening on %s\n", srv.URL)

	select {}
}
