package social

import "strings"

// Tag represents a single Tag
type Tag string

// String converts the user into a string, prepending a "#"
func (t Tag) String() string {
	return "#" + string(t)
}

// Tags represents multiple tags
type Tags []Tag

// AsStrings converts the tags into a slice of String
func (t Tags) AsStrings() (tags []string) {
	tags = make([]string, len(t))
	for i, tag := range t {
		tags[i] = tag.String()
	}
	return
}

// String converts the tags into a comma separated list of Tag
func (t Tags) String() (str string) {
	var strTags = t.AsStrings()
	str = strings.Join(strTags, ",")
	return
}
