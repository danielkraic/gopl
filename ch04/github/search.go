package github

import (
	"net/url"
	"strings"
	"net/http"
	"fmt"
	"encoding/json"
)

func SearchIssues(terms []string) (*IssuesSearchResult, error) {
	q := url.QueryEscape(strings.Join(terms, " "))
	resp, err := http.Get(IssueUrl + "?q=" + q)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("search query failed with code %d\n", resp.StatusCode)
	}

	var result IssuesSearchResult
	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		return nil, fmt.Errorf("Invalif json api response. Error: %s\n", resp.StatusCode)
	}
	return &result, nil
}
