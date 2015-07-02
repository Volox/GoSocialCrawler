package social

import "time"
import "fmt"

// Post represents a post in a social network
type Post struct {
	ID       string
	Text     string
	Date     time.Time
	Location *Location
	Author   string
	AuthorID string
	Tags     Tags
	Lang     string
	Source   string
	Nil      int
}

// String converts the post to a string representation
func (p Post) String() (str string) {
	str = fmt.Sprintf("%s from %s[%s] {%s}", p.ID, p.Author, p.AuthorID, p.Tags)
	return
}

// Posts represents a list of posts
type Posts []Post

// Location represents a GeoJSON point in the format [ Longitude, Latitude ]
type Location [2]float64

// Latitude returns the latitude of the Location
func (l *Location) Latitude() float64 {
	return l[1]
}

// Longitude returns the longitude of the Location
func (l *Location) Longitude() float64 {
	return l[0]
}

// NewLocation return a valid Location struct based on lat ang long info
func NewLocation(lat float64, lon float64) *Location {
	return &Location{lon, lat}
}
