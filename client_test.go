package complyadvantage

import (
	"testing"
	"net/http"
	"fmt"
	"os"
)

const (
	TestSearchTerm = "vladimir putin"
	BirthYearTest  = "1952"
)

var DefaultClient = NewClient("OedyQopDdU41uOtVBc8t7kbgvCnO0tHu") // todo env

func TestClient_Searches(t *testing.T) {
	r, err := DefaultClient.Searches(&SearchesRequest{
		SearchTerm: TestSearchTerm,
		Filters: &Filter{
			BirthYear: BirthYearTest,
		},
	})
	if err != nil {
		t.Fatal(err)
	}

	if r.Code != http.StatusOK {
		t.Fatalf("status code is %d, error message: %s", r.Code, r.Message)
	}

	if r.Content.Data.SearchTerm != TestSearchTerm {
		t.Fatalf("%s (resp) != %s (req)", r.Content.Data.SearchTerm, TestSearchTerm)
	}
}

func TestClient_SearchesByTerm(t *testing.T) {
	r, err := DefaultClient.SearchesByTerm(TestSearchTerm)
	if err != nil {
		t.Fatal(err)
	}

	if r.Code != http.StatusOK {
		t.Fatalf("status code is %d", r.Code)
	}

	if r.Content.Data.SearchTerm != TestSearchTerm {
		t.Fatalf("%s (resp) != %s (req)", r.Content.Data.SearchTerm, TestSearchTerm)
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
	sr, err := DefaultClient.SearchesByTerm(TestSearchTerm)
	if err != nil {
		t.Fatal(err)
	}

	if sr.Code != http.StatusOK {
		t.Fatalf("status code is %d", sr.Code)
	}

	if sr.Content.Data.SearchTerm != TestSearchTerm {
		t.Fatalf("%s (resp) != %s (req)", sr.Content.Data.SearchTerm, TestSearchTerm)
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
	sr, err := DefaultClient.SearchesByTerm(TestSearchTerm)
	if err != nil {
		t.Fatal(err)
	}

	if sr.Code != http.StatusOK {
		t.Fatalf("status code is %d", sr.Code)
	}

	if sr.Content.Data.SearchTerm != TestSearchTerm {
		t.Fatalf("%s (resp) != %s (req)", sr.Content.Data.SearchTerm, TestSearchTerm)
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

func TestClient_SearchesByIdDetails(t *testing.T) {
	sr, err := DefaultClient.SearchesByTerm(TestSearchTerm)
	if err != nil {
		t.Fatal(err)
	}

	if sr.Code != http.StatusOK {
		t.Fatalf("status code is %d", sr.Code)
	}

	if sr.Content.Data.SearchTerm != TestSearchTerm {
		t.Fatalf("%s (resp) != %s (req)", sr.Content.Data.SearchTerm, TestSearchTerm)
	}

	sir, err := DefaultClient.SearchesByIdDetails(sr.Content.Data.Id)
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
