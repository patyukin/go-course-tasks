package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	for i := 0; i < 1000; i++ {
		req, err := http.NewRequest("GET", "http://localhost:7071", nil)
		if err != nil {
			fmt.Errorf("asdfasdf")
		}

		c := http.Client{}
		res, err := c.Do(req)
		if err != nil {
			return
		}

		r, _ := io.ReadAll(res.Body)

		fmt.Print(string(r) + "\r\n")
	}
}
