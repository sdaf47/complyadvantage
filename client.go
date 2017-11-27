package complyadvantage

import (
	"net/http"
	"io/ioutil"
	"encoding/json"
	"io"
	"bytes"
	"fmt"
)

const (
	Users = "/users"

	Searches             = "/searches"
	SearchesById         = "/searches/%d"
	SearchesByIdCert     = "/searches/%d/certificate"
	SearchesByIdDetails  = "/searches/%d/details"
	SearchesByIdMonitors = "/searches/%d/monitors"
)

const (
	BaseUrl = "https://api.complyadvantage.com"
	ApiKey  = "api_key"
)

type Client struct {
	httpClient *http.Client
	apiKey     string
}

func NewClient(apiKey string) *Client {
	httpClient := http.DefaultClient
	c := &Client{
		httpClient: httpClient,
		apiKey:     apiKey,
	}

	return c
}

func (c *Client) send(method, url string, body io.Reader) (response []byte, err error) {
	r, err := http.NewRequest(method, BaseUrl+url, body)
	if err != nil {
		return nil, err
	}

	query := r.URL.Query()
	// todo per_page
	// todo page
	// todo sort_by
	// todo sort_dir

	query.Add(ApiKey, c.apiKey)
	r.URL.RawQuery = query.Encode()

	result, err := c.httpClient.Do(r)
	if err != nil {
		return nil, err
	}

	buffer, err := ioutil.ReadAll(result.Body)
	if err != nil {
		return nil, err
	}

	return buffer, nil
}

func (c *Client) get(url string) ([]byte, error) {
	return c.send(http.MethodGet, url, nil)
}

func (c *Client) post(url string, request Request) ([]byte, error) {
	br, err := json.Marshal(request)
	if err != nil {
		panic(err)
	}

	return c.send(http.MethodPost, url, bytes.NewReader(br))
}

func (c *Client) jsonPost(url string, request Request, r Response) (Response, error) {
	data, err := c.post(url, request)
	if err != nil {
		return nil, err
	}

	return readJson(data, r)
}

func (c *Client) jsonGet(url string, r Response) (Response, error) {
	data, err := c.get(url)
	if err != nil {
		return nil, err
	}

	return readJson(data, r)
}

func (c *Client) pdfGet(url string) ([]byte, error) {
	return c.get(url)
}

func (c *Client) Searches(request *SearchesRequest) (*SearchResponse, error) {
	r, err := c.jsonPost(Searches, request, &SearchResponse{})

	return r.(*SearchResponse), err
}

func (c *Client) Users() (*UsersResponse, error) {
	r, err := c.jsonGet(Users, &UsersResponse{})

	return r.(*UsersResponse), err
}

func (c *Client) SearchesByTerm(term string) (r *SearchResponse, err error) {
	request := &SearchesRequest{
		SearchTerm: term,
	}

	return c.Searches(request)
}

func (c *Client) SearchesById(id uint) (*SearchResponse, error) {
	url := fmt.Sprintf(SearchesById, id)
	r, err := c.jsonGet(url, &SearchResponse{})

	return r.(*SearchResponse), err
}

func (c *Client) SearchesByIdCert(id uint) ([]byte, error) {
	url := fmt.Sprintf(SearchesByIdCert, id)

	return c.pdfGet(url)
}

func (c *Client) SearchesByIdDetails(id uint) (*SearchResponse, error) {
	url := fmt.Sprintf(SearchesByIdDetails, id)
	r, err := c.jsonGet(url, &SearchResponse{})

	return r.(*SearchResponse), err
}
