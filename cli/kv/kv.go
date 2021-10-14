package kv

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/couchbase/gocb/v2"
	"github.com/spf13/cobra"
)

func getTranscoder(transcoder string) gocb.Transcoder {
	if transcoder == "json" {
		return gocb.NewJSONTranscoder()
	} else if transcoder == "raw_json" {
		return gocb.NewRawJSONTranscoder()
	}
	return gocb.NewRawStringTranscoder()
}

func durationWrapper(f func() (interface{}, error)) (interface{}, time.Duration, error) {
	begin := time.Now()
	result, err := f()
	end := time.Since(begin)
	return result, end, err
}

func encodeValue(value string, transcoder string) (interface{}, error) {
	if transcoder == "json" {
		valueI := map[string]interface{}{}
		err := json.Unmarshal([]byte(value), &valueI)
		if err != nil {
			return nil, err
		}
		return valueI, nil
	}
	return value, nil
}

func KvCmdBuilder() *cobra.Command {
	var kvCmd = &cobra.Command{Use: "kv"}
	setDefaultFlags(kvCmd)

	var cmdGet = &cobra.Command{
		Use:   "get",
		Short: "Get values by keys",
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) < 1 {
				fmt.Println("Pass keys in args")
				return
			}
			collection := getCurCollection(dsn, bucket, username, password, scope, collection)
			for _, key := range args {
				fmt.Printf("Getting key: %s\n", args)
				getResultI, elapsed, err := durationWrapper(
					func() (interface{}, error) {
						return collection.Get(
							key,
							&gocb.GetOptions{
								Timeout:    timeout,
								Transcoder: getTranscoder(transcoder),
							},
						)
					},
				)
				if err != nil {
					fmt.Printf("Problem with receiving key %s: %s\n", key, err)
					continue
				}
				getResult := getResultI.(*gocb.GetResult)
				fmt.Printf("value received in %s\n", elapsed)
				if transcoder == "json" {
					var result map[string]interface{}
					if err := getResult.Content(&result); err == nil {
						buf, err := json.Marshal(result)
						if err != nil {
							fmt.Println("Problem with marshaling data to json")
						}
						fmt.Printf("Key: %s\nValue: %s\n", key, string(buf))
					} else {
						fmt.Printf("Problem with encoding: %s\n", err)
					}
				} else {
					var result string
					if err := getResult.Content(&result); err == nil {
						fmt.Printf("Key: %s\nValue: %s\n", key, result)
					} else {
						fmt.Printf("Problem with encoding: %s\n", err)
					}
				}

			}
		},
	}
	setDefaultFlags((cmdGet))
	kvCmd.AddCommand(cmdGet)

	{
		var ttl time.Duration
		var value string
		var key string
		setCommandFlags := func(cmd *cobra.Command) {
			setDefaultFlags(cmd)
			cmd.Flags().DurationVarP(&ttl, "ttl", "", 15*time.Minute, "TTL: 1m, 1s, 100ms, etc. Default 10m")
			cmd.Flags().StringVarP(&key, "key", "k", "", "Key")
			cmd.Flags().StringVarP(&value, "value", "v", "", "Value")
		}
		var cmdUpsert = &cobra.Command{
			Use:   "upsert",
			Short: "Upsert value for key",
			Run: func(cmd *cobra.Command, args []string) {
				fmt.Printf("Upserting key %s by value \"%s\"\n", key, value)
				collection := getCurCollection(dsn, bucket, username, password, scope, collection)
				opts := gocb.UpsertOptions{
					Expiry:     ttl,
					Timeout:    timeout,
					Transcoder: getTranscoder(transcoder),
				}
				valueI, err := encodeValue(value, transcoder)
				if err != nil {
					return
				}
				_, err = collection.Upsert(
					key, valueI,
					&opts,
				)
				if err != nil {
					fmt.Printf("Problem with upserting value: %s\n", err)
					return
				}
				fmt.Println("Upserted!")
			},
		}
		setCommandFlags(cmdUpsert)
		kvCmd.AddCommand(cmdUpsert)

		var cmdInsert = &cobra.Command{
			Use:   "insert",
			Short: "Insert value for key",
			Run: func(cmd *cobra.Command, args []string) {
				fmt.Printf("Inserting key %s value \"%s\"\n", key, value)
				collection := getCurCollection(dsn, bucket, username, password, scope, collection)
				opts := gocb.InsertOptions{
					Expiry:     ttl,
					Timeout:    timeout,
					Transcoder: getTranscoder(transcoder),
				}
				valueI, err := encodeValue(value, transcoder)
				if err != nil {
					return
				}
				_, err = collection.Insert(
					key, valueI,
					&opts,
				)
				if err != nil {
					return
				}
				fmt.Println("Saved!")
			},
		}
		setCommandFlags(cmdInsert)
		kvCmd.AddCommand(cmdInsert)

		var cmdReplace = &cobra.Command{
			Use:   "replace",
			Short: "Replace value for key",
			Run: func(cmd *cobra.Command, args []string) {
				fmt.Printf("Replacing key %s value by value \"%s\"\n", key, value)
				collection := getCurCollection(dsn, bucket, username, password, scope, collection)
				opts := gocb.ReplaceOptions{
					Expiry:     ttl,
					Timeout:    timeout,
					Transcoder: getTranscoder(transcoder),
				}
				valueI, err := encodeValue(value, transcoder)
				if err != nil {
					return
				}
				_, err = collection.Replace(
					key, valueI,
					&opts,
				)
				if err != nil {
					fmt.Printf("Problem with replacing value: %s\n", err)
					return
				}
				fmt.Println("Saved!")
			},
		}
		setCommandFlags(cmdReplace)
		kvCmd.AddCommand(cmdReplace)

		var cmdRemove = &cobra.Command{
			Use:   "remove",
			Short: "remove keys",
			Run: func(cmd *cobra.Command, args []string) {
				if len(args) < 1 {
					if len(args) < 1 {
						fmt.Println("Pass keys in args")
						return
					}
				}
				fmt.Printf("Removing keys: %s\n", args)
				collection := getCurCollection(dsn, bucket, username, password, scope, collection)
				opts := gocb.RemoveOptions{
					Timeout: timeout,
				}
				for _, key := range args {
					_, err := collection.Remove(
						key,
						&opts,
					)
					if err != nil {
						fmt.Printf("Problem with removing key: %s\n", err)
						return
					}
				}
				fmt.Println("Removed!")
			},
		}
		setCommandFlags(cmdRemove)
		kvCmd.AddCommand(cmdRemove)
	}

	return kvCmd
}
