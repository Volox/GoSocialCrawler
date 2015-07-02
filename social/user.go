package social

import "strings"

// User represent a single user
type User string

// String converts the user into a string, prepending a "@"
func (u User) String() string {
	return "@" + string(u)
}

// Users represent multiple users
type Users []User

// AsStrings converts the tags into a slice of String
func (u Users) AsStrings() (users []string) {
	users = make([]string, len(u))
	for i, user := range u {
		users[i] = user.String()
	}
	return
}

// String converts the tags into a comma separated list of Tag
func (u Users) String() (str string) {
	var strUsers = u.AsStrings()
	str = strings.Join(strUsers, ",")
	return
}
