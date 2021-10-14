package kv

import (
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/couchbase/gocb/v2"
	"github.com/spf13/cobra"
)

func BucketCmdBuilder() *cobra.Command {
	var bucketCmd = &cobra.Command{Use: "bucket"}
	setDefaultFlags(bucketCmd)

	var pause time.Duration
	var num int
	var cmdPing = &cobra.Command{
		Use:   "ping",
		Short: "Ping nodes",
		Run: func(cmd *cobra.Command, args []string) {
			cbBucket := getCurBucket(dsn, bucket, username, password)
			fmt.Printf("Pinging bucket %s\n", bucket)
			var lastPingResult *gocb.PingResult
			for i := 1; i <= num; i++ {
				pingResult, err := cbBucket.Ping(&gocb.PingOptions{
					ReportID:     "my-report",
					ServiceTypes: []gocb.ServiceType{gocb.ServiceTypeKeyValue},
					Timeout:      timeout,
				})
				if err != nil {
					panic(err)
				}
				lastPingResult = pingResult
				for service, pingReports := range pingResult.Services {
					if service != gocb.ServiceTypeKeyValue {
						panic("we got a service type that we didn't ask for!")
					}

					for _, pingReport := range pingReports {
						if pingReport.State != gocb.PingStateOk {
							fmt.Printf(
								"Node %s at remote %s is not OK, error: %s, latency: %s\n",
								pingReport.ID, pingReport.Remote, pingReport.Error, pingReport.Latency.String(),
							)
						} else {
							fmt.Printf(
								"Node %s at remote %s is OK, latency: %s\n",
								pingReport.ID, pingReport.Remote, pingReport.Latency.String(),
							)
						}
					}
				}
				time.Sleep(pause)
			}

			b, err := json.Marshal(lastPingResult)
			if err != nil {
				panic(err)
			}

			fmt.Printf("Ping report JSON: %s", string(b))
		},
	}
	setDefaultFlags(cmdPing)
	cmdPing.Flags().DurationVarP(&pause, "pause", "", 10*time.Second, "Pause: 1m, 1s, 100ms, etc. Default 10s")
	cmdPing.Flags().IntVarP(&num, "num", "n", 10, "Ping: 1m, 1s, 100ms, etc. Default 10.")
	bucketCmd.AddCommand(cmdPing)

	var cmdDiagnostics = &cobra.Command{
		Use:   "diagnostics",
		Short: "Run node diagnostics",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Printf("Pinging bucket %s\n", bucket)
			cluster, err := getCluster(dsn, username, password)
			if err != nil {
				panic(err)
			}
			_, err = getBucket(cluster, bucket)
			if err != nil {
				panic(err)
			}
			diagnostics, err := cluster.Diagnostics(&gocb.DiagnosticsOptions{
				ReportID: "my-report",
			})
			if err != nil {
				panic(err)
			}

			if diagnostics.State != gocb.ClusterStateOnline {
				log.Printf("Overall cluster state is not online\n")
			} else {
				log.Printf("Overall cluster state is online\n")
			}

			for serviceType, diagReports := range diagnostics.Services {
				for _, diagReport := range diagReports {
					if diagReport.State != gocb.EndpointStateConnected {
						fmt.Printf(
							"Node %s at remote %s is not connected on service %s, activity last seen at: %s\n",
							diagReport.ID, diagReport.Remote, serviceType, diagReport.LastActivity.String(),
						)
					} else {
						fmt.Printf(
							"Node %s at remote %s is connected on service %s, activity last seen at: %s\n",
							diagReport.ID, diagReport.Remote, serviceType, diagReport.LastActivity.String(),
						)
					}
				}
			}

			db, err := json.Marshal(diagnostics)
			if err != nil {
				panic(err)
			}

			fmt.Printf("Diagnostics report JSON: %s", string(db))
		},
	}
	setDefaultFlags((cmdDiagnostics))
	bucketCmd.AddCommand(cmdDiagnostics)

	return bucketCmd
}
