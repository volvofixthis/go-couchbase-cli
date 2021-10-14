package main

import (
	"bitbucket.rbc.ru/go/go-couchbase-cli/cli/kv"
	"github.com/spf13/cobra"
)

func main() {
	var rootCmd = &cobra.Command{Use: "app"}
	rootCmd.AddCommand(kv.KvCmdBuilder())
	rootCmd.AddCommand(kv.BucketCmdBuilder())
	rootCmd.Execute()
}
