package complyadvantage

import (
	"testing"
	"net/http"
	"fmt"
	"os"
)

const (
	SearchTerm = "Robert Gabriel Mugabe"
)

var DefaultClient = NewClient("OedyQopDdU41uOtVBc8t7kbgvCnO0tHu") // todo env

func TestClient_Searches(t *testing.T) {
	r, err := DefaultClient.SearchesByTerm(SearchTerm)
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

func TestClient_SearchesById(t *testing.T) {
	sr, err := DefaultClient.SearchesByTerm(SearchTerm)
	if err != nil {
		t.Fatal(err)
	}

	if sr.Code != http.StatusOK {
		t.Fatalf("status code is %d", sr.Code)
	}

	if sr.Content.Data.SearchTerm != SearchTerm {
		t.Fatalf("%s (resp) != %s (req)", sr.Content.Data.SearchTerm, SearchTerm)
	}

	sir, err := DefaultClient.SearchesById(sr.Content.Data.Id)
	if err != nil {
		t.Fatal(err)
	}

	if sir.Code != http.StatusOK {
		t.Fatalf("status code is %d", sir.Code)
	}

	if sir.Content.Data.SearchTerm != sr.Content.Data.SearchTerm {
		t.Fatalf("status code is %d", sir.Code)
	}
}

func TestClient_SearchesByIdCert(t *testing.T) {
	sr, err := DefaultClient.SearchesByTerm(SearchTerm)
	if err != nil {
		t.Fatal(err)
	}

	if sr.Code != http.StatusOK {
		t.Fatalf("status code is %d", sr.Code)
	}

	if sr.Content.Data.SearchTerm != SearchTerm {
		t.Fatalf("%s (resp) != %s (req)", sr.Content.Data.SearchTerm, SearchTerm)
	}

	sir, err := DefaultClient.SearchesByIdCert(sr.Content.Data.Id)
	if err != nil {
		t.Fatal(err)
	}

	pdfFile, err := os.Create(fmt.Sprintf("searches_certificate_%d.pdf", sr.Content.Data.Id))
	if err != nil {
		t.Fatal("cant create file", err)
	}

	_, err = pdfFile.Write(sir)
	if err != nil {
		t.Fatal("cant write to file", err)
	}

	pdfFile.Close()
}
