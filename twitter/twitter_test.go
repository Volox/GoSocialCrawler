package twitter

import (
	"fmt"
	"github.com/Volox/GoCrawler/social"
	"os"
	"testing"
)

var credentials = &social.Credentials{
	Key:    "sCCiDYz9iNtm2KVw4jHAA",
	Secret: "M8oCAm1KNd04RlPrkmGRGjoWvYUcIqY8my8eS3txcjQ",
}
var tw *Twitter

func TestMain(m *testing.M) {
	tw, _ = New(credentials)
	os.Exit(m.Run())
}

func printResults(posts social.Posts) {
	for _, post := range posts {
		fmt.Printf("Got post: %v\n", post)
	}
}

func TestTag(t *testing.T) {
	var tag = social.Tag("EXPO")
	_, err := tw.GetByTag(tag, nil, nil)
	if err != nil {
		t.Error(err)
	}
}

func TestTags(t *testing.T) {
	var tags = social.Tags{"EXPO", "Milano"}
	_, err := tw.GetByTags(tags, nil, nil)
	if err != nil {
		t.Error(err)
	}
}

func TestUser(t *testing.T) {
	var user = social.User("expo2015milano")
	_, err := tw.GetByUser(user, nil, nil)
	if err != nil {
		t.Error(err)
	}
}

func TestUsers(t *testing.T) {
	var users = social.Users{"expo2015milano", "comunemi"}
	_, err := tw.GetByUsers(users, nil, nil)
	if err != nil {
		t.Error(err)
	}
}

func TestPagination(t *testing.T) {
	var tag = social.Tag("EXPO")

	_, err := tw.GetByTag(tag, nil, nil)
	if err != nil {
		t.Error(err)
	}
}
