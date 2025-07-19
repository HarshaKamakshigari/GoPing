package main

import (
	"fmt"
	"github.com/HarshaKamakshigari/GoPing/checker"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: goping <url1> <url2> ...")
		return
	}

	urls := os.Args[1:]

	results := checker.CheckWebsites(urls)

	checker.PrintResults(results)
}
