package main

import (
	"net/http"

	"github.com/edzhabs/social/internal/store"
	"github.com/edzhabs/social/utils"
)

type CommentPayload struct {
	Content string `json:"content" validate:"required,max=1000"`
}

func (app *application) postCommentHandler(w http.ResponseWriter, r *http.Request) {
	var payload CommentPayload

	if err := utils.ReadJSON(w, r, &payload); err != nil {
		app.internalServerError(w, r, err)
		return
	}

	if err := utils.Validate.Struct(payload); err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	post := app.getPostFromCtx(r)

	user := app.getUserFromCtx(r)

	comment := &store.Comment{
		PostID:  post.ID,
		Content: payload.Content,
		UserID:  user.ID,
	}

	ctx := r.Context()

	if err := app.store.Comments.Create(ctx, comment); err != nil {
		app.internalServerError(w, r, err)
		return
	}

	if err := utils.ResponseJSON(w, http.StatusCreated, comment); err != nil {
		app.internalServerError(w, r, err)
		return
	}
}
