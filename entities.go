package complyadvantage

import "time"

// 1 or 0
type JsonBool int

type SearchProfile struct {
	Name string `json:"name"`
	Slug string `json:"slug"`
}

type Field struct {
	Locale string `json:"locale"`
	Name   string `json:"name"`
	Values string `json:"values"`
	Tag    string `json:"tag"`
	Source string `json:"source"`
}

type AkaNames struct {
	Name string `json:"name"`
}

type Associate struct {
	Name        string `json:"name"`
	Association string `json:"association"`
}

type Tags struct {
	Name interface{}
}

type Media struct {
	Date    time.Time `json:"date"`
	PdfUrl  string    `json:"pdf_url"`
	Snippet string    `json:"snippet"`
	Title   string    `json:"title"`
	Url     string    `json:"url"`
}

type Entity struct {
	Id            string       `json:"id"`
	EntityType    string       `json:"entity_type"`
	LastUpdateUtc time.Time    `json:"last_update_utc"`
	Name          string       `json:"name"`
	FirstName     string       `json:"first_name"`
	MiddleName    string       `json:"middle_name"`
	LastName      string       `json:"last_name"`
	Sources       []string     `json:"sources"`
	Types         []string     `json:"types"`
	Fields        []*Field     `json:"fields"`
	Aka           []*AkaNames  `json:"aka"`
	Associates    []*Associate `json:"associates"`
	Media         []*Media     `json:"media"`
	IsWhitelisted bool         `json:"is_whitelisted"`
	MatchTypes    []string     `json:"match_types"`
	Score         float64      `json:"score"`
}

type User struct {
	Id        uint   `json:"id"`
	Email     string `json:"email"`
	Name      string `json:"name"`
	Phone     string `json:"phone"`
	UpdatedAt string `json:"updated_at"`
	CreatedAt string `json:"created_at"`
}

type Filter struct {
	Types     []string `json:"types"`
	BirthYear string   `json:"birth_year"`

	// 1 or 0
	RemoveDeceased int `json:"remove_deceased"`

	Passport   string  `json:"passport"`
	EntityType string  `json:"entity_type"`
	ExactMatch bool    `json:"exact_match"`
	Fuzziness  float64 `json:"fuzziness"`
}

type Hit struct {
	Doc Entity `json:"doc"`
}
