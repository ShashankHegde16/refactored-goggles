package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	resp, err := http.Get("http://google.com")

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	bs := make([]byte, 32*1024)
	resp.Body.Read(bs)
	fmt.Println(string(bs))

}
