package checker

import (
	"crypto/tls"
	"fmt"
	"net"
	"net/http"
	"strings"
	"sync"
	"time"
)

type Result struct {
	URL          string
	Status       string
	ResponseTime int64
	SSLEnd       string
	LocalIP      string
	RemoteIP     string
}

func CheckWebsites(urls []string) []Result {
	var wg sync.WaitGroup
	results := make([]Result, len(urls))

	for i, url := range urls {
		wg.Add(1)
		go func(i int, url string) {
			defer wg.Done()
			results[i] = check(url)
		}(i, url)
	}

	wg.Wait()
	return results
}
