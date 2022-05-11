package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

func main() {
	for _, url := range os.Args[1:] {
		resp, err := http.Get(url)
		if err != nil {
			return
		}
		b, err := ioutil.ReadAll(resp.Body)
		state := resp.StatusCode
		resp.Body.Close()
		if err != nil {
			return
		}

		fmt.Printf("%s", b)
		fmt.Printf("状态码为%d", state)
	}
}
