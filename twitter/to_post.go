package twitter

import "github.com/Volox/GoCrawler/social"
import "github.com/kurrik/twittergo"

// ToPost converts a tweet into a Post structure
func (tw Twitter) ToPost(data interface{}) (post social.Post, err error) {
	// Fixed
	post.Source = "twitter"

	var tweet = data.(twittergo.Tweet)
	post.ID = tweet.IdStr()
	post.Text = tweet.Text()
	post.Date = tweet.CreatedAt()

	// Language
	post.Lang = tweet.Language()

	// Tags
	var tags = make(social.Tags, 0)
	if tweet["entities"] != nil {
		var entities = tweet["entities"].(map[string]interface{})
		if entities["hashtags"] != nil {
			var hashtags = entities["hashtags"].([]interface{})
			for _, item := range hashtags {
				var hashtag = item.(map[string]interface{})
				var tag = social.Tag(hashtag["text"].(string))
				tags = append(tags, tag)
			}
		}
	}
	post.Tags = tags

	// Author
	var user = tweet.User()
	post.Author = user.Name()
	post.AuthorID = user.IdStr()

	// Location
	var location *social.Location
	if tweet["coordinates"] != nil {
		var geoPoint = tweet["coordinates"].(map[string]interface{})
		var coordinates = geoPoint["coordinates"].([]interface{})
		var lon = coordinates[0].(float64)
		var lat = coordinates[1].(float64)
		location = social.NewLocation(lat, lon)
	}
	post.Location = location

	// Nil
	post.Nil = -1

	return post, nil
}
