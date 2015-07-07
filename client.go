package rabbitmq

import (
	"bytes"
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"runtime"
	"strings"
)

const (
	Version = "0.1.0"
)

var UserAgent string

func init() {
	UserAgent = "rabbitmq-golang/" + Version + " (" + runtime.GOOS + "-" + runtime.GOARCH + ")"
}

type ClientOptFunc func(*Client) error

type Client struct {
	c   *http.Client
	url string
}

func NewClient(options ...ClientOptFunc) (*Client, error) {
	c := &Client{
		c:   http.DefaultClient,
		url: "http://localhost:15672",
	}

	if url := os.Getenv("RABBITMQ_URL"); url != "" {
		c.url = url
	}

	for _, opt := range options {
		if err := opt(c); err != nil {
			return nil, err
		}
	}

	// Normalize URL (http://127.0.0.1:9200/path?query=1 -> http://127.0.0.1:9200)
	u, err := url.Parse(c.url)
	if err != nil {
		return nil, err
	}
	u.Fragment = ""
	u.Path = ""
	u.RawQuery = ""
	c.url = u.String()

	return c, nil
}

func SetHttpClient(httpClient *http.Client) ClientOptFunc {
	return func(c *Client) error {
		if httpClient != nil {
			c.c = httpClient
		} else {
			c.c = http.DefaultClient
		}
		return nil
	}
}

func SetURL(urlStr string) ClientOptFunc {
	return func(c *Client) error {
		u, err := url.Parse(urlStr)
		if err != nil {
			return err
		}
		c.url = u.String()
		return nil
	}
}

func (c *Client) Execute(method, path string, params url.Values, body interface{}) (*Response, error) {
	var query string
	if len(params) > 0 {
		query = "?" + params.Encode()
	}

	req, err := newRequest(method, c.url+path+query)
	if err != nil {
		return nil, err
	}

	if body != nil {
		switch b := body.(type) {
		case string:
			req.SetBodyString(b)
		default:
			req.SetBodyJSON(b)
		}
	}

	res, err := c.c.Do((*http.Request)(req))
	if err != nil {
		return nil, err
	}
	if res.Body != nil {
		defer res.Body.Close()
	}

	if err := checkResponse(res); err != nil {
		return nil, err
	}

	return newResponse(res)
}

// -- Request --

type Request http.Request

func newRequest(method, url string) (*Request, error) {
	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Add("User-Agent", UserAgent)
	req.Header.Add("Accept", "application/json")
	return (*Request)(req), nil
}

func (r *Request) SetBodyJSON(data interface{}) error {
	body, err := json.Marshal(data)
	if err != nil {
		return err
	}
	r.SetBody(bytes.NewReader(body))
	r.Header.Set("Content-Type", "application/json")
	return nil
}

func (r *Request) SetBodyString(body string) error {
	return r.SetBody(strings.NewReader(body))
}

func (r *Request) SetBody(body io.Reader) error {
	rc, ok := body.(io.ReadCloser)
	if !ok && body != nil {
		rc = ioutil.NopCloser(body)
	}
	r.Body = rc
	if body != nil {
		switch v := body.(type) {
		case *strings.Reader:
			r.ContentLength = int64(v.Len())
		case *bytes.Buffer:
			r.ContentLength = int64(v.Len())
		}
	}
	return nil
}

// -- Response --

type Response struct {
	StatusCode int
	Header     http.Header
	Body       json.RawMessage
}

func newResponse(res *http.Response) (*Response, error) {
	r := &Response{
		StatusCode: res.StatusCode,
		Header:     res.Header,
	}
	if res.Body != nil {
		slurp, err := ioutil.ReadAll(res.Body)
		if err != nil {
			return nil, err
		}
		if len(slurp) > 0 {
			if err := json.Unmarshal(slurp, &r.Body); err != nil {
				return nil, err
			}
		}
	}
	return r, nil
}
