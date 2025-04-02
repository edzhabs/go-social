package main

import (
	"net/http"

	"github.com/edzhabs/social/internal/store"
	"github.com/edzhabs/social/utils"
)

func (app *application) getUserFeedHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	fq := store.PaginatedQuery{
		Limit:  20,
		Offset: 0,
		Sort:   "desc",
	}

	fq, err := fq.Parse(r)
	if err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	if err := utils.Validate.Struct(fq); err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	//TODO: change the userid during auth
	feed, err := app.store.Posts.GetUserFeed(ctx, 655, fq)
	if err != nil {
		app.internalServerError(w, r, err)
		return
	}

	if err := utils.ResponseJSON(w, http.StatusOK, feed); err != nil {
		app.internalServerError(w, r, err)
		return
	}
}
