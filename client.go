package tlsg10x

import (
	"fmt"
	"io"
	"net/http"
	"net/http/cookiejar"
	"regexp"
)

type Client struct {
	Api
	Host     string
	Username string
	Password string

	client *http.Client
}

// new tlsg108e client
func New(h, u, p string, httpClient *http.Client) *Client {
	if httpClient == nil {
		httpClient = http.DefaultClient
		// required by v6 hardware versions
		jar, _ := cookiejar.New(nil)
		httpClient.Jar = jar
	}

	return &Client{
		Host:     h,
		Username: u,
		Password: p,
		client:   httpClient,
	}
}

func (t *Client) getPage(url string) (string, error) {
	err := t.Login()
	if err != nil {
		return "", fmt.Errorf("login error: invalid username or password (maybe host?)")
	}

	resp, err := t.client.Get(url)
	if err != nil {
		return "", fmt.Errorf("error retrieving information from url %s", url)
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("error retrieving information from url %s", url)
	}

	err = t.Logout()
	if err != nil {
		return "", fmt.Errorf("logout error")
	}
	return string(body), nil
}

func (t *Client) getBytes(url string) ([]byte, error) {
	err := t.Login()
	if err != nil {
		return nil, err
	}

	resp, err := t.client.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	err = t.Logout()
	if err != nil {
		return nil, err
	}

	return body, nil
}

func (t *Client) getFirstValue(str string, regex string) (string, error) {
	r := regexp.MustCompile(regex)
	findResult := r.FindStringSubmatch(str)
	if len(findResult) <= 1 {
		return "", fmt.Errorf("not found")
	}

	return findResult[1], nil
}
