// Package github provides a Go API for the GitHub issue tracker.
// See https://developer.github.com/v3/search/#search-issues.

//包 github 为 GitHub 问题跟踪器提供了一个 Go API
package github

import "time"

const IssuesURL = "https://api.github.com/search/issues"

type IssuesSearchResult struct {
	TotalCount int `json:"total_count"`
	Items      []*Issue
}

type Issue struct {
	Number    int
	HTMLURL   string `json:"html_url"`
	Title     string
	State     string
	User      *user
	CreatedAt time.Time `json:"created_at"`
	Body      string    // in Markdown format
}

type user struct {
	Login   string
	HTMLURL string `json:"html_url"`
}
