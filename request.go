package complyadvantage

type Request interface{}

type SearchTerm interface{}

type SearchesRequest struct {
	// required
	SearchTerm    SearchTerm        `json:"search_term"`
	ClientRef     string            `json:"client_ref"`
	SearchProfile string            `json:"search_profile"`
	Fuzziness     float64           `json:"fuzziness"`
	Offset        int               `json:"offset"`
	Limit         int               `json:"limit"`
	Filters       *Filter           `json:"filters"`
	Tags          map[string]string `json:"tags"`
}

type SearchString string

type SearchObject struct {
	LastName   string `json:"last_name"`
	MiddleName string `json:"middle_name"`
	FirstName  string `json:"first_name"`
}
