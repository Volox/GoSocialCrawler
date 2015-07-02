package main

import (
	log "github.com/Sirupsen/logrus"
	// "os"
	// "github.com/Volox/GoCrawler"
	"github.com/Volox/GoCrawler/social"
	"github.com/Volox/GoCrawler/twitter"
)

func init() {
	log.SetLevel(log.DebugLevel)
}

func main() {
	var tw *twitter.Twitter

	var credentials = &social.Credentials{
		Key:    "sCCiDYz9iNtm2KVw4jHAA",
		Secret: "M8oCAm1KNd04RlPrkmGRGjoWvYUcIqY8my8eS3txcjQ",
	}

	tw, _ = twitter.New(credentials)

	var tag = social.Tag("Maserati")

	var options = social.DefaultOptions()
	options.Count = 1000
	options.MaxPages = 0
	// options.Paginate = false

	posts, pages, err := tw.GetByTag(tag, options, nil)
	if err != nil {
		log.Error(err)
	}
	log.Infof("Got %d posts from %d page/s", len(posts), pages)
}
