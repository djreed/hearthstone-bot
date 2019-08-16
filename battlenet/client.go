package battlenet

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

const (
	LIB_VERSION     = "0.1"
	USER_AGENT      = "go-http/" + LIB_VERSION
	BASE_URL_FORMAT = "https://%s.api.blizzard.com/"
)

// Battlenet API client
type Client struct {
	// Authed HTTP client for requests
	Client *http.Client

	// API base URL
	BaseURL *url.URL

	// Client's HTTP UserAgent
	UserAgent string
}

func NewClient(region string, c *http.Client) *Client {
	region = strings.ToLower(region)

	if c == nil {
		c = http.DefaultClient
	}

	baseURLStr := fmt.Sprintf(BASE_URL_FORMAT, region)

	baseURL, err := url.Parse(baseURLStr)
	if err != nil {
		// We panic because we manually construct it above so it should
		// never really fail unless the user gives us a REALLY bad region.
		panic(err)
	}

	return &Client{
		Client:    c,
		BaseURL:   baseURL,
		UserAgent: USER_AGENT,
	}
}

func (c *Client) Hearthstone() *HearthService {
	return &HearthService{client: c}
}

func (c *Client) NewRequest(method, urlStr string, body interface{}) (*http.Request, error) {
	rel, err := url.Parse(urlStr)
	if err != nil {
		return nil, err
	}

	u := c.BaseURL.ResolveReference(rel)

	var buf io.ReadWriter
	if body != nil {
		buf = new(bytes.Buffer)
		err := json.NewEncoder(buf).Encode(body)
		if err != nil {
			return nil, err
		}
	}

	req, err := http.NewRequest(method, u.String(), buf)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Accept", "application/json")
	if c.UserAgent != "" {
		req.Header.Add("User-Agent", c.UserAgent)
	}

	return req, nil
}

func (c *Client) Do(req *http.Request, v interface{}) (*Response, error) {
	resp, err := c.Client.Do(req)
	if err != nil {
		return nil, err
	}

	defer func() {
		// Drain up to 512 bytes and close the body to let the Transport reuse the connection
		io.CopyN(ioutil.Discard, resp.Body, 512)
		resp.Body.Close()
	}()

	response := newResponse(resp)

	if err := CheckError(resp); err != nil {
		return response, err
	}

	if v != nil {
		if w, ok := v.(io.Writer); ok {
			io.Copy(w, resp.Body)
		} else {
			err = json.NewDecoder(resp.Body).Decode(v)
			if err == io.EOF {
				err = nil // ignore EOF errors caused by empty response body
			}
		}
	}

	return response, err
}
