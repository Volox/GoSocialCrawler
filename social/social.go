package social

// Social represents a generic SocialNetwork
type Social interface {
	Run(runID string, runType string, options *Options, status *Status) (err error)
	UpdateStatus() (err error)

	ToPost(data interface{}) (post Post, err error)

	// Get data methods
	GetByTag(tag Tag, options *Options, status *Status) (posts Posts, pages uint, err error)
	GetByTags(tags Tags, options *Options, status *Status) (posts Posts, pages uint, err error)
	GetByUser(user User, options *Options, status *Status) (posts Posts, pages uint, err error)
	GetByUsers(users Users, options *Options, status *Status) (posts Posts, pages uint, err error)
	GetByPoint(point Point, options *Options, status *Status) (posts Posts, pages uint, err error)
	GetByPoints(points Points, options *Options, status *Status) (posts Posts, pages uint, err error)
}
