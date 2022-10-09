package api

import "net/http"

func (a *API) PullRequestComment(w http.ResponseWriter, r *http.Request) {
	writeResponse(w, http.StatusOK, "ok")
}
