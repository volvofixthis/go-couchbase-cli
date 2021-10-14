package kv

import (
	"time"

	"github.com/spf13/cobra"
)

var (
	username   string
	password   string
	dsn        string
	bucket     string
	scope      string
	collection string
	timeout    time.Duration
	transcoder string
)

func setDefaultFlags(cmd *cobra.Command) {
	cmd.Flags().StringVarP(&dsn, "dsn", "d", "", "Cluster DSN")
	cmd.Flags().StringVarP(&bucket, "bucket", "b", "", "Bucket")
	cmd.Flags().StringVarP(&username, "username", "u", "", "Username")
	cmd.Flags().StringVarP(&password, "password", "p", "", "Password")
	cmd.Flags().StringVarP(&scope, "scope", "s", "", "Scope")
	cmd.Flags().StringVarP(&collection, "collection", "c", "_default", "Collection")
	cmd.Flags().DurationVarP(&timeout, "timeout", "t", 10*time.Second, "Timeout: 1m, 1s, 100ms, etc.")
	cmd.Flags().StringVarP(&transcoder, "transcoder", "", "string", "Transcoder: json, raw_json, string")
}
