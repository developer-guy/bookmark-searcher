package main

import (
	"github.com/developer-guy/bookmarks-searcher/parser"
	"github.com/developer-guy/bookmarks-searcher/pub"
	"fmt"
	"github.com/developer-guy/bookmarks-searcher/bookmark"
	"github.com/developer-guy/bookmarks-searcher/search"
	"context"
	"github.com/developer-guy/bookmarks-searcher/env"
	"time"
	"strconv"
)

type message struct {
	Date_Added               string `json:"date_added"`
	Date_Modified            string `json:"date_modified"`
	Id                       string `json:"id"`
	Name                     string `json:"name"`
	Sync_Transaction_Version string `json:"sync_transaction_version"`
	Type                     string `json:"type"`
}

const REDIS_CHANNEL = "bookmark"

func main() {
	ctx := context.Background()
	var (
		redisCloudPassword = env.GetEnvironmentOrDefault("REDIS_CLOUD_PASSWORD", "")
		redisCloudHost     = env.GetEnvironmentOrDefault("REDIS_CLOUD_HOST", "")
		redisCloudDb       = env.GetEnvironmentOrDefault("REDIS_CLOUD_DB", "")
	)

	var (
		elasticHost = env.GetEnvironmentOrDefault("ELASTIC_HOST", "")
	)

	elastic := &search.ElasticOpt{
		ElasticHost: elasticHost,
	}

	elasticClient, err := elastic.Client()

	if err != nil {
		// Handle error
		panic(err)
	}

	clusterInfo, err := elasticClient.ClusterHealth().Do(ctx)
	if err != nil {
		// Handle error
		panic(err)
	}
	fmt.Printf("Node name :%s Status: %s", clusterInfo.ClusterName, clusterInfo.Status)

	item, err := parser.Parse(".")

	if err != nil {
		panic(err)
	}

	redis := &pub.RedisOpt{
		Address:    redisCloudHost,
		Password:   redisCloudPassword,
		DB:         redisCloudDb,
		MaxRetries: 2,
	}

	client := redis.Connect()

	var bookMarks bookmark.Bookmarks = item.(bookmark.Bookmarks)
	counter := 0

	for _, v := range bookMarks.Roots.BookmarkBar.Children {
		for _, b := range v.Children {
			dateAddedMillisecond, _ := strconv.ParseInt(b.Date_Added, 10, 64)
			message := b.Name + "::" + b.Url + "::" + time.Unix(0, dateAddedMillisecond*100).Format("02-01-2006 15:04:05")
			timer := time.NewTimer(time.Second * 2)
			_, err := client.Publish(REDIS_CHANNEL, message).Result()
			<-timer.C
			if err != nil {
				panic(err)
			}
			fmt.Println(message, "published from redis in", REDIS_CHANNEL, "channel.")
			counter++
		}
	}

	defer client.Close()
	fmt.Println(counter, "number of records published...")

}
