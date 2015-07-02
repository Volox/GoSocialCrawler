package twitter

import (
	"fmt"
	log "github.com/Sirupsen/logrus"
	"github.com/Volox/GoCrawler/social"
	"github.com/kurrik/oauth1a"
	"github.com/kurrik/twittergo"
	"net/http"
	"strconv"
)

// Twitter implements the Social interface and represent the Twitter implementation
type Twitter struct {
	api         *twittergo.Client
	credentials *social.Credentials
}

// New creates a new instance of the Twitter Social
func New(credentials *social.Credentials) (twitter *Twitter, err error) {

	log.Debugln("Creating a new Twitter Social")
	var user *oauth1a.UserConfig

	if !credentials.AppOnly() {
		log.Debugln("Application only authorization")
		user = oauth1a.NewAuthorizedConfig(credentials.Token, credentials.TokenSecret)
	}

	config := &oauth1a.ClientConfig{
		ConsumerKey:    credentials.Key,
		ConsumerSecret: credentials.Secret,
	}

	client := twittergo.NewClient(config, user)

	// Assign for return value
	twitter = &Twitter{
		api:         client,
		credentials: credentials,
	}
	return
}

// Get is the generic method used to perform a call to the REST APIs
func (tw Twitter) Get(query Query, options *social.Options) (posts social.Posts, pages uint, err error) {
	var req *http.Request
	var resp *twittergo.APIResponse

	log.Debugln("Performing a request")

	// Init options
	if options == nil {
		log.Debugln("No option passed, using default ones")
		options = social.DefaultOptions()
	}
	if options.Paginate == false {
		options.MaxPages = 1
	}
	log.WithFields(options.ToFields()).Debugf("Options")

	// Init query
	query.Set("count", strconv.FormatUint(uint64(options.Count), 10))
	log.WithFields(query.ToFields()).Debugf("Query")

	posts = make(social.Posts, 0)
	for pages < options.MaxPages || options.MaxPages == 0 {

		// Create the URL
		url := fmt.Sprintf("/1.1/search/tweets.json?%v", query.Encode())
		// Create the HTTP Request
		req, err = http.NewRequest("GET", url, nil)
		if err != nil {
			return
		}
		// Send the Auth request
		resp, err = tw.api.SendRequest(req)
		if err != nil {
			return
		}
		// Increment the total number of pages retrieved
		pages++

		// Parse results
		var results = &twittergo.SearchResults{}
		err = resp.Parse(results)
		if err != nil {
			return
		}

		// Convert data to posts
		for _, tweet := range results.Statuses() {
			post, _ := tw.ToPost(tweet)
			posts = append(posts, post)
		}

		log.Debugf("Got page %d with %d tweets", pages, len(results.Statuses()))

		// Handle pagination, chack if a new page is needed
		values, nextQueryError := results.NextQuery()
		if nextQueryError != nil {
			log.Debug(nextQueryError)
			break
		}
		query = Query{values}
		log.WithFields(query.ToFields()).Debugf("New page query")
	}

	log.Debugf("Query ended, retrieved %d posts", len(posts))
	return
}
