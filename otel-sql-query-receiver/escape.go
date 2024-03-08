package main

import (
	"fmt"
	"net/url"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run escape.go <password>")
		return
	}

	fmt.Println(url.QueryEscape(os.Args[1]))
}
