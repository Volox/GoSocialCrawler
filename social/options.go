package social

import (
	"fmt"
	log "github.com/Sirupsen/logrus"
)

// Options to pass to the method to run
type Options struct {
	Paginate bool
	Count    uint
	MaxPages uint
}

// DefaultOptions returns a default set of options
func DefaultOptions() *Options {
	var options = &Options{}

	options.Count = 100
	options.Paginate = true
	options.MaxPages = 50

	return options
}

// ToFields converts the options into the log.Fields type supported by the logger
func (o *Options) ToFields() (fields log.Fields) {
	fields = make(log.Fields)
	fields["paginate"] = o.Paginate
	fields["count"] = o.Count
	fields["maxpages"] = o.MaxPages
	return
}

// String implements the Stringer interface so it can be used in fmt.Printf and similar
func (o *Options) String() string {
	var str string

	str += fmt.Sprintf("paginate=%v", o.Paginate)
	str += fmt.Sprintf("Count: %v", o.Count)
	str += fmt.Sprintf("MaxPages: %v", o.MaxPages)

	return str
}
