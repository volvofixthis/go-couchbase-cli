package kv

import (
	"errors"
	"time"

	"github.com/couchbase/gocb/v2"
)

func getCluster(dsn, username, password string) (*gocb.Cluster, error) {
	if dsn == "" {
		return nil, errors.New("Dsn is required")
	}
	opts := gocb.ClusterOptions{
		Username: username,
		Password: password,
		TimeoutsConfig: gocb.TimeoutsConfig{
			ConnectTimeout:   time.Second * 2,
			KVTimeout:        time.Millisecond * 30,
			KVDurableTimeout: time.Millisecond * 100,
		},
	}
	cluster, err := gocb.Connect(
		dsn,
		opts,
	)
	if err != nil {
		return nil, err
	}
	return cluster, nil
}

func getBucket(cluster *gocb.Cluster, bucket string) (*gocb.Bucket, error) {
	if bucket == "" {
		return nil, errors.New("bucket is required")
	}
	cbBucket := cluster.Bucket(bucket)

	err := cbBucket.WaitUntilReady(5*time.Second, nil)
	if err != nil {
		return nil, err
	}
	return cbBucket, nil
}

func getCollection(bucket *gocb.Bucket, scope, collection string) (*gocb.Collection, error) {
	var cbCollection *gocb.Collection
	if scope != "" {
		scope := bucket.Scope(scope)
		if collection != "" {
			cbCollection = scope.Collection(collection)
		} else {
			cbCollection = scope.Collection("_default")
		}
	} else {
		cbCollection = bucket.Collection(collection)
	}
	return cbCollection, nil
}

func getCurCollection(dsn, bucket, username, password, scope, collection string) *gocb.Collection {
	cluster, err := getCluster(dsn, username, password)
	if err != nil {
		panic(err)
	}
	cbBucket, err := getBucket(cluster, bucket)
	if err != nil {
		panic(err)
	}
	cbCollection, err := getCollection(cbBucket, scope, collection)
	if err != nil {
		panic(err)
	}
	return cbCollection
}

func getCurBucket(dsn, bucket, username, password string) *gocb.Bucket {
	cluster, err := getCluster(dsn, username, password)
	if err != nil {
		panic(err)
	}
	cbBucket, err := getBucket(cluster, bucket)
	if err != nil {
		panic(err)
	}
	return cbBucket
}
