package slack

import (
	"io/ioutil"
	"net/http"
	"net/url"
	"path"
	"strings"

	"github.com/google/go-querystring/query"
)

const apiURL = "https://slack.com/api/"

type Client struct {
	APIClient APIClient

	Chat *ChatMethod
}

func NewClient() *Client {
	apiClient := newAPIClient()

	return &Client{
		APIClient: apiClient,

		Chat: newChatMethod(apiClient),
	}
}

type APIClient interface {
	Post(uri string, values interface{}) ([]byte, int, error)
}

type apiClient struct {
	HttpClient *http.Client
}

func newAPIClient() *apiClient {
	client := &http.Client{}

	return &apiClient{
		HttpClient: client,
	}
}

func (apiClient *apiClient) Post(uri string, values interface{}) ([]byte, int, error) {
	q, err := query.Values(values)
	if err != nil {
		return nil, 0, err
	}

	u, err := url.Parse(apiURL)
	if err != nil {
		return nil, 0, err
	}

	u.Path = path.Join(u.Path, uri)

	resp, err := apiClient.HttpClient.Post(u.String(), "application/x-www-form-urlencoded", strings.NewReader(q.Encode()))
	if err != nil {
		return nil, 0, err
	}
	defer func() {
		_ = resp.Body.Close()
	}()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, resp.StatusCode, err
	}

	return body, resp.StatusCode, nil
}
