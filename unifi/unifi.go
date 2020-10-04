package unifi //

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"path"
	"strings"
	"sync"
)

const (
	apiPath    = "/api"
	apiPathNew = "/proxy/network/api"

	loginPath    = "/api/login"
	loginPathNew = "/api/auth/login"
)

type NotFoundError struct{}

func (err *NotFoundError) Error() string {
	return "not found"
}

type APIError struct {
	RC      string
	Message string
}

func (err *APIError) Error() string {
	return err.Message
}

type Client struct {
	// single thread client calls for CSRF, etc.
	sync.Mutex

	c       *http.Client
	baseURL *url.URL

	apiPath   string
	loginPath string

	csrf string
}

func (c *Client) CSRFToken() string {
	return c.csrf
}

func (c *Client) SetBaseURL(base string) error {
	var err error
	c.baseURL, err = url.Parse(base)
	if err != nil {
		return err
	}

	// error for people who are still passing hard coded old paths
	if path := strings.TrimSuffix(c.baseURL.Path, "/"); path == apiPath {
		return fmt.Errorf("expected a base URL without the `/api`, got: %q", c.baseURL)
	}

	return nil
}

func (c *Client) SetHTTPClient(hc *http.Client) error {
	c.c = hc
	return nil
}

func (c *Client) setAPIUrlStyle(ctx context.Context) error {
	// check if new style API
	// this is modified from the unifi-poller (https://github.com/unifi-poller/unifi) implementation.
	// see https://github.com/unifi-poller/unifi/blob/4dc44f11f61a2e08bf7ec5b20c71d5bced837b5d/unifi.go#L101-L104
	// and https://github.com/unifi-poller/unifi/commit/43a6b225031a28f2b358f52d03a7217c7b524143

	req, err := http.NewRequestWithContext(ctx, "GET", c.baseURL.String(), nil)
	if err != nil {
		return err
	}

	// We can't share these cookies with other requests, so make a new client.
	// Checking the return code on the first request so don't follow a redirect.
	client := &http.Client{
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			return http.ErrUseLastResponse
		},
		Transport: c.c.Transport,
	}

	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	_, _ = io.Copy(ioutil.Discard, resp.Body)

	if resp.StatusCode == http.StatusOK {
		// the new API returns a 200 for a / request
		c.apiPath = apiPathNew
		c.loginPath = loginPathNew
		return nil
	}

	// The old version returns a "302" (to /manage) for a / request
	c.apiPath = apiPath
	c.loginPath = loginPath
	return nil
}

func (c *Client) Login(ctx context.Context, user, pass string) error {
	if c.c == nil {
		c.c = &http.Client{}

		jar, _ := cookiejar.New(nil)
		c.c.Jar = jar
	}

	err := c.setAPIUrlStyle(ctx)
	if err != nil {
		return fmt.Errorf("unable to determine API URL style: %w", err)
	}

	err = c.do(ctx, "POST", c.loginPath, &struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}{
		Username: user,
		Password: pass,
	}, nil)
	if err != nil {
		return err
	}

	return nil
}

func (c *Client) do(ctx context.Context, method, relativeURL string, reqBody interface{}, respBody interface{}) error {
	// single threading requests, this is mostly to assist in CSRF token propagation
	c.Lock()
	defer c.Unlock()

	var (
		reqReader io.Reader
		err       error
		reqBytes  []byte
	)
	if reqBody != nil {
		reqBytes, err = json.Marshal(reqBody)
		if err != nil {
			return fmt.Errorf("unable to marshal JSON: %s %s %w", method, relativeURL, err)
		}
		reqReader = bytes.NewReader(reqBytes)
	}

	reqURL, err := url.Parse(relativeURL)
	if err != nil {
		return fmt.Errorf("unable to parse URL: %s %s %w", method, relativeURL, err)
	}
	if !strings.HasPrefix(relativeURL, "/") && !reqURL.IsAbs() {
		reqURL.Path = path.Join(c.apiPath, reqURL.Path)
	}

	url := c.baseURL.ResolveReference(reqURL)
	req, err := http.NewRequestWithContext(ctx, method, url.String(), reqReader)
	if err != nil {
		return fmt.Errorf("unable to create request: %s %s %w", method, relativeURL, err)
	}

	req.Header.Set("User-Agent", "terraform-provider-unifi/0.1")
	req.Header.Add("Content-Type", "application/json; charset=utf-8")

	if c.csrf != "" {
		req.Header.Set("X-CSRF-Token", c.csrf)
	}

	resp, err := c.c.Do(req)
	if err != nil {
		return fmt.Errorf("unable to perform request: %s %s %w", method, relativeURL, err)
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusNotFound {
		return &NotFoundError{}
	}

	if csrf := resp.Header.Get("x-csrf-token"); csrf != "" {
		c.csrf = resp.Header.Get("x-csrf-token")
	}

	if resp.StatusCode != 200 {
		fmt.Printf("Request Body:\n%s\n", string(reqBytes))
		errBody := struct {
			Meta meta `json:"meta"`
		}{}
		err = json.NewDecoder(resp.Body).Decode(&errBody)
		return fmt.Errorf("%s %s (%s) for %s %s", errBody.Meta.RC, errBody.Meta.Message, resp.Status, method, url.String())
	}

	if respBody == nil || resp.ContentLength == 0 {
		return nil
	}

	// TODO: check rc in addition to status code?

	err = json.NewDecoder(resp.Body).Decode(respBody)
	if err != nil {
		return fmt.Errorf("unable to decode body: %s %s %w", method, relativeURL, err)
	}

	return nil
}

type meta struct {
	RC      string `json:"rc"`
	Message string `json:"msg"`
}

func (m *meta) error() error {
	if m.RC != "ok" {
		return &APIError{
			RC:      m.RC,
			Message: m.Message,
		}
	}

	return nil
}
