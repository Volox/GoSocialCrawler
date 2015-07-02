package twitter

import (
	log "github.com/Sirupsen/logrus"

	"net/url"
)
import "fmt"

// Query represent the key-value pairs sent as query parameters in the GET request
type Query struct {
	url.Values
}

// String implements the Stringer interface so it can be used in fmt.Printf and similar
func (q Query) String() (str string) {
	for key, val := range q.Values {
		str += fmt.Sprintf("%s = %v", key, val)
	}
	return
}

// ToFields converts the query data into the log.Fields type supported by the logger
func (q *Query) ToFields() (fields log.Fields) {
	fields = make(log.Fields)
	for key, val := range q.Values {
		fields[key] = val
	}
	return
}

// NewQuery create and returns an usable Query struct
func NewQuery() (q Query) {
	q = Query{}
	q.Values = url.Values{}
	return
}
