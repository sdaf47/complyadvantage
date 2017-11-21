package complyadvantage

import (
	"testing"
	"net/http"
)

const (
	SearchTerm = "Robert Gabriel Mugabe"
)

var DefaultClient = NewClient("OedyQopDdU41uOtVBc8t7kbgvCnO0tHu") // todo env

func TestClient_Get(t *testing.T) {
	resp := &UsersResponse{}

	err := DefaultClient.get(Users, resp)
	if err != nil {
		t.Fatalf("%s", err)
	}
}

func TestClient_Post(t *testing.T) {
	resp := &SearchResponse{}

	req := &SearchesRequest{
		SearchTerm: SearchTerm,
		Fuzziness:  0.7,
	}

	err := DefaultClient.post(Searches, resp, req)
	if err != nil {
		t.Fatal("error: ", err)
	}

	if resp.Content.Data.SearchTerm != SearchTerm {
		t.Fatalf("%s (resp) != %s (req)", resp.Content.Data.SearchTerm, SearchTerm)
	}
}

func TestClient_Searches(t *testing.T) {
	r, err := DefaultClient.Searches(SearchTerm)
	if err != nil {
		t.Fatal(err)
	}

	if r.Code != http.StatusOK {
		t.Fatalf("status code is %d", r.Code)
	}

	if r.Content.Data.SearchTerm != SearchTerm {
		t.Fatalf("%s (resp) != %s (req)", r.Content.Data.SearchTerm, SearchTerm)
	}
}

func TestClient_Users(t *testing.T) {
	r, err := DefaultClient.Users()

	if err != nil {
		t.Fatal(err)
	}

	if r.Code != http.StatusOK {
		t.Fatalf("status code is %d", r.Code)
	}
}
