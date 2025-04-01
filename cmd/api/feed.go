package main

import (
	"net/http"

	"github.com/edzhabs/social/utils"
)

func (app *application) getUserFeedHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	feed, err := app.store.Posts.GetUserFeed(ctx, 655)
	if err != nil {
		app.internalServerError(w, r, err)
		return
	}

	if err := utils.ResponseJSON(w, http.StatusOK, feed); err != nil {
		app.internalServerError(w, r, err)
		return
	}
}
