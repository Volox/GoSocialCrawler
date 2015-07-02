package twitter

import "github.com/Volox/GoCrawler/social"
import "strings"

// GetByTag implments the twitter specific behaviour of the Social method GetByTag
func (tw Twitter) GetByTag(tag social.Tag, options *social.Options, status *social.Status) (social.Posts, uint, error) {
	var q = NewQuery()
	q.Set("q", tag.String())
	return tw.Get(q, options)
}

// GetByTags implments the twitter specific behaviour of the Social method GetByTags
func (tw Twitter) GetByTags(tags social.Tags, options *social.Options, status *social.Status) (social.Posts, uint, error) {
	var q = NewQuery()
	var strTags = tags.AsStrings()
	q.Set("q", strings.Join(strTags, " OR "))
	return tw.Get(q, options)
}

// GetByUser implments the twitter specific behaviour of the Social method GetByUser
func (tw Twitter) GetByUser(user social.User, options *social.Options, status *social.Status) (social.Posts, uint, error) {
	var q = NewQuery()
	q.Set("q", user.String())
	return tw.Get(q, options)
}

// GetByUsers implments the twitter specific behaviour of the Social method GetByUsers
func (tw Twitter) GetByUsers(users social.Users, options *social.Options, status *social.Status) (social.Posts, uint, error) {
	var q = NewQuery()
	var strUsers = users.AsStrings()
	q.Set("q", strings.Join(strUsers, " OR "))
	return tw.Get(q, options)
}

// GetByPoint implments the twitter specific behaviour of the Social method GetByPoint
func (tw Twitter) GetByPoint(point social.Point, options *social.Options, status *social.Status) (social.Posts, uint, error) {
	return nil, 0, nil
}

// GetByPoints implments the twitter specific behaviour of the Social method GetByPoints
func (tw Twitter) GetByPoints(points social.Points, options *social.Options, status *social.Status) (social.Posts, uint, error) {
	return nil, 0, nil
}
