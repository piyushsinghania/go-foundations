package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {
	resp, err := http.Get("http://formester.com")

	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	// ioutil.WriteFile(filename, []byte(d.toString()), 0666)
	io.Copy(os.Stdout, resp.Body)
}
