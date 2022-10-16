package api

import (
	"net/http"
	"strconv"

	"github.com/google/go-github/v45/github"
)

const (
	owner = "ryanrolds"
)

type pullRequestComment struct {
	owner   string
	repo    string
	issue   int
	comment string
}

// curl -H "X-Access-Token=..." http://localhost:8080/pull_request/comment?repo=...&comment=...&issue=...
func (a *API) PullRequestComment(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	// Parse and validate request details
	comment, err := pullRequestCommentFromRequest(r)
	if err != nil {
		writeResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	// Create requested comment on pull request
	_, _, err = a.ghAppClient.Issues.CreateComment(ctx, comment.owner, comment.repo, comment.issue, &github.IssueComment{
		Body: &comment.comment,
	})
	if err != nil {
		writeResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	writeResponse(w, http.StatusOK, "ok")
}

func pullRequestCommentFromRequest(r *http.Request) (*pullRequestComment, error) {
	qs := r.URL.Query()

	repo := qs.Get("repo")
	comment := qs.Get("comment")

	issueRaw := qs.Get("issue")
	issue, err := strconv.Atoi(issueRaw)
	if err != nil {
		return nil, err
	}

	prc := pullRequestComment{
		owner:   owner,
		repo:    repo,
		issue:   issue,
		comment: comment,
	}

	return &prc, nil
}
