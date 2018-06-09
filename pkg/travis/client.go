package travis

import (
	"net/http"
	"time"
)

const apiURL = "https://api.travis-ci.org/"

// Option is a functional option for configuring the API client
type Option func(*Client) error

// BaseURL allows overriding of API client baseURL for testing
func BaseURL(baseURL string) Option {
	return func(c *Client) error {
		c.baseURL = baseURL
		return nil
	}
}

func RepoSlug(repoSlug string) Option {
	return func(c *Client) error {
		c.repoSlug = repoSlug
		return nil
	}
}

func ApiToken(apiToken string) Option {
	return func(c *Client) error {
		c.apiToken = apiToken
		return nil
	}
}

// parseOptions parses the supplied options functions and returns a configured
// *Client instance
func (c *Client) parseOptions(opts ...Option) error {
	// Range over each options function and apply it to our API type to
	// configure it. Options functions are applied in order, with any
	// conflicting options overriding earlier calls.
	for _, option := range opts {
		err := option(c)
		if err != nil {
			return err
		}
	}

	return nil
}

// Client holds information necessary to make a request to your API
type Client struct {
	baseURL        string
	repoSlug       string
	apiToken       string
	httpClient     *http.Client
}

// New creates a new API client
func New(opts ...Option) (*Client, error) {
	client := &Client{
		baseURL:  apiURL,
		httpClient: &http.Client{
			Timeout: time.Second * 30,
		},
	}

	if err := client.parseOptions(opts...); err != nil {
		return nil, err
	}

	return client, nil
}
