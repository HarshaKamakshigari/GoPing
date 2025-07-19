package checker

import (
	"fmt"
)

func PrintResults(results []Result) {
	fmt.Println("==========================================================================================================================")
	fmt.Printf("%-30s %-10s %-10s %-12s %-30s\n", "URL", "STATUS", "TIME(ms)", "SSL EXPIRY", "LOCAL → REMOTE IP")
	fmt.Println("------------------------------------------------------------------------------------")

	for _, r := range results {
		fmt.Printf(
			"%-30s %-10s %-10d %-12s %-30s\n",
			r.URL,
			r.Status,
			r.ResponseTime,
			r.SSLEnd,
			fmt.Sprintf("%s → %s", r.LocalIP, r.RemoteIP),
		)
	}
	fmt.Println("==========================================================================================================================")
}
