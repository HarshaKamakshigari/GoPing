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

func check(url string) Result {
	start := time.Now()
	result := Result{URL: url}

	client := http.Client{Timeout: 5 * time.Second}
	resp, err := client.Get(url)
	if err != nil {
		result.Status = "DOWN"
	} else {
		result.Status = fmt.Sprintf("%d %s", resp.StatusCode, http.StatusText(resp.StatusCode))
		result.ResponseTime = time.Since(start).Milliseconds()
		resp.Body.Close()
	}

	host := stripHTTPS(url)

	if ssl, err := getSSLExpiry(host); err == nil {
		result.SSLEnd = ssl
	} else {
		result.SSLEnd = "N/A"
	}

	if local, remote, err := getConnectionInfo(host); err == nil {
		result.LocalIP = local
		result.RemoteIP = remote
	} else {
		result.LocalIP = "?"
		result.RemoteIP = "?"
	}

	return result
}

func getSSLExpiry(host string) (string, error) {
	conn, err := tls.Dial("tcp", host+":443", nil)
	if err != nil {
		return "", err
	}
	defer conn.Close()

	cert := conn.ConnectionState().PeerCertificates[0]
	daysLeft := int(time.Until(cert.NotAfter).Hours() / 24)
	return fmt.Sprintf("%d days", daysLeft), nil
}

func getConnectionInfo(host string) (string, string, error) {
	conn, err := net.DialTimeout("tcp", host+":80", 3*time.Second)
	if err != nil {
		return "", "", err
	}
	defer conn.Close()
	return conn.LocalAddr().String(), conn.RemoteAddr().String(), nil
}

func stripHTTPS(url string) string {
	url = strings.TrimPrefix(url, "https://")
	url = strings.TrimPrefix(url, "http://")
	return strings.Split(url, "/")[0]
}
