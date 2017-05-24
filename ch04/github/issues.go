package github

import (
	"time"
)

const IssueUrl = "https://api.github.com/search/issues"

type IssuesSearchResult struct {
	TotalCount int `json:"total_count"`
	Items []*Issue
}

type Issue struct {
	Number int
	HTMLUrl string `json:"html_url"`
	Title string
	State string
	User *User
	CreatedAt time.Time `json:"create_at"`
	Body string
}

type User struct {
	Login string
	HTMLUrl string `json:"html_url"`
}
