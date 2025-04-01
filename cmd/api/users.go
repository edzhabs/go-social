package main

import (
	"context"
	"errors"
	"net/http"
	"strconv"

	"github.com/edzhabs/social/internal/store"
	"github.com/edzhabs/social/utils"
	"github.com/go-chi/chi/v5"
)

type userKey string

var userCtx userKey = "user"

func (app *application) getUserHandler(w http.ResponseWriter, r *http.Request) {
	user := app.getUserFromCtx(r)

	if err := utils.ResponseJSON(w, http.StatusOK, user); err != nil {
		app.internalServerError(w, r, err)
		return
	}
}

type FollowUser struct {
	UserID int64 `json:"user_id" validate:"required"`
}

func (app *application) followUserHandler(w http.ResponseWriter, r *http.Request) {
	followedUser := app.getUserFromCtx(r)

	// TODO: revert back to auth from ctx
	var payload FollowUser

	if err := utils.ReadJSON(w, r, &payload); err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	if followedUser.ID == payload.UserID {
		app.badRequestResponse(w, r, errors.New("cannot follow own user_id"))
		return
	}

	if err := app.store.Followers.Follow(r.Context(), followedUser.ID, payload.UserID); err != nil {
		switch err {
		case store.ErrConflict:
			app.conflictResponse(w, r, err)
		default:
			app.internalServerError(w, r, err)
		}
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func (app *application) unfollowUserHandler(w http.ResponseWriter, r *http.Request) {
	followedUser := app.getUserFromCtx(r)

	// TODO: revert back to auth from ctx
	var payload FollowUser

	if err := utils.ReadJSON(w, r, &payload); err != nil {
		app.internalServerError(w, r, err)
		return
	}

	if err := app.store.Followers.Unfollow(r.Context(), followedUser.ID, payload.UserID); err != nil {
		app.internalServerError(w, r, err)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func (app *application) userContextMiddle(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		userID, err := strconv.ParseInt(chi.URLParam(r, "userID"), 10, 64)
		if err != nil {
			app.internalServerError(w, r, err)
			return
		}

		ctx := r.Context()

		user, err := app.store.Users.GetByID(ctx, userID)
		if err != nil {
			switch err {
			case store.ErrNotFound:
				app.notFoundResponse(w, r, err)
			default:
				app.internalServerError(w, r, err)
			}
			return
		}

		ctx = context.WithValue(ctx, userCtx, user)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func (app *application) getUserFromCtx(r *http.Request) *store.User {
	user, _ := r.Context().Value(userCtx).(*store.User)
	return user
}
