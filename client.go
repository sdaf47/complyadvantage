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
	BaseUrl     = "https://api.complyadvantage.com"
	ContentType = "application/json"
	ApiKey      = "api_key"
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

func (c *Client) send(method, url string, response Response, body io.Reader) (error) {
	r, err := http.NewRequest(method, BaseUrl+url, body)
	if err != nil {
		return err
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
		return err
	}

	buffer, err := ioutil.ReadAll(result.Body)
	if err != nil {
		return err
	}

	if err = json.Unmarshal(buffer, response); err != nil {
		return err
	}

	return nil
}

func (c *Client) get(url string, response Response) (error) {
	return c.send(http.MethodGet, url, response, nil)
}

func (c *Client) post(url string, response Response, request interface{}) (error) {
	br, err := json.Marshal(request)
	if err != nil {
		panic(err)
	}

	return c.send(http.MethodPost, url, response, bytes.NewReader(br))
}

func (c *Client) Search(req *SearchesRequest) (r *SearchResponse, err error) {
	r = &SearchResponse{}
	if err := c.post(Searches, r, req); err != nil {
		return nil, err
	}

	return r, err
}

func (c *Client) Users() (r *UsersResponse, err error) {
	r = &UsersResponse{}
	if err := c.get(Users, r); err != nil {
		return nil, err
	}

	return r, nil
}

func (c *Client) Searches(term string) (r *SearchResponse, err error) {
	req := &SearchesRequest{
		SearchTerm: term,
	}

	return c.Search(req)
}

func (c *Client) SearchesById(id uint) (r *SearchResponse, err error) {
	r = &SearchResponse{}
	url := fmt.Sprintf(SearchesById, id)
	if err := c.get(url, r); err != nil {
		return nil, err
	}

	return r, nil
}

func (c *Client) SearchesByIdCert(id uint) (r *BaseResponse, err error) {
	// todo
	r = &BaseResponse{}
	url := fmt.Sprintf(SearchesByIdCert, id)
	if err := c.get(url, r); err != nil {
		return nil, err
	}

	return r, nil
}
