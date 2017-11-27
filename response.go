package complyadvantage

import (
	"encoding/json"
)

const (
	EntityTypePerson  = "person"
	EntityTypeCompany = "company"
	/// todo ...
)

type Response interface{}

type Content struct {
	Data interface{} `json:"data"`
}

type BaseResponse struct {
	Code    int     `json:"code"`
	Status  string  `json:"status"`
	Content Content `json:"content"`
	Message string  `json:"message"`
}

type UsersResponse struct {
	BaseResponse
	Content struct {
		Data []User `json:"data"`
	} `json:"content"`
}

type SearchResult struct {
	Id            uint           `json:"id"`
	Ref           string         `json:"ref"`
	SearcherId    uint           `json:"searcher_id"`
	Searcher      Searcher       `json:"assignee"`
	Assignee      Searcher       `json:"assignee"`
	AssigneeId    uint           `json:"assignee_id"`
	SearchProfile *SearchProfile `json:"search_profile"`
	Filters       *Filter        `json:"filters"`
	MatchStatus   string         `json:"match_status"`
	RiskLevel     string         `json:"risk_level"`
	SearchTerm    string         `json:"search_term"`
	TotalHits     uint           `json:"total_hits"`
	UpdatedAt     string         `json:"updated_at"`
	CreatedAt     string         `json:"created_at"`
	Tags          *Tags          `json:"tags"`
	Hits          []Hit          `json:"hits"`
	ShareUrl      string         `json:"share_url"`
}

type SearchResponse struct {
	BaseResponse
	Content struct {
		Data SearchResult `json:"data"`
	} `json:"content"`
}

type SearchesResponse struct {
	BaseResponse
	Content struct {
		Data []SearchResult `json:"data"`
	} `json:"content"`
}

type CertificateResponse struct {
	PdfData []byte
}

func readJson(data []byte, v Response) (Response, error) {
	err := json.Unmarshal(data, v)

	return v, err
}
