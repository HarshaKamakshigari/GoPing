package checker

import (
	"fmt"
	"github.com/olekukonko/tablewriter"
	"os"
)

func PrintResults(results []Result) {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"URL", "Status", "Time (ms)", "SSL Expiry", "Local → Remote IP"})

	for _, r := range results {
		table.Append([]string{
			r.URL,
			r.Status,
			fmt.Sprintf("%d", r.ResponseTime),
			r.SSLEnd,
			fmt.Sprintf("%s → %s", r.LocalIP, r.RemoteIP),
		})
	}

	table.Render()
}
