package complyadvantage

type SearchesRequest struct {
	// required
	SearchTerm    string            `json:"search_term"`
	ClientRef     string            `json:"client_ref"`
	SearchProfile string            `json:"search_profile"`
	Fuzziness     float64           `json:"fuzziness"`
	Offset        int               `json:"offset"`
	Limit         int               `json:"limit"`
	Filters       *Filter           `json:"filters"`
	Tags          map[string]string `json:"tags"`
}
