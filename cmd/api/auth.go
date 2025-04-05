package main

import (
	"net/http"

	"github.com/edzhabs/social/internal/store"
	"github.com/edzhabs/social/utils"
	"github.com/google/uuid"
)

type RegisterUserPayload struct {
	Username        string `json:"username" validate:"required,alphanum,max=100"`
	Email           string `json:"email" validate:"required,email,max=255"`
	Password        string `json:"password" validate:"required,min=3,max=100"`
	ConfirmPassword string `json:"confirm_password" validate:"required,eqfield=Password"`
}

type UserWithToken struct {
	*store.User
	Token string `json:"token"`
}

func (app *application) registerUserHandler(w http.ResponseWriter, r *http.Request) {
	var payload RegisterUserPayload

	if err := utils.ReadJSON(w, r, &payload); err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	if err := utils.Validate.Struct(payload); err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	user := &store.User{
		Username: payload.Username,
		Email:    payload.Email,
	}

	if err := user.Password.Set(payload.Password); err != nil {
		app.internalServerError(w, r, err)
		return
	}

	ctx := r.Context()

	plainToken := uuid.New().String()
	hashToken := utils.HashToken(plainToken)

	if err := app.store.Users.CreateAndInvite(ctx, user, hashToken, app.config.mail.expiry); err != nil {
		switch err {
		case store.ErrDuplicateEmail, store.ErrDuplicateUsername:
			app.badRequestResponse(w, r, err)
		default:
			app.internalServerError(w, r, err)
		}
		return
	}

	userWithToken := UserWithToken{
		User:  user,
		Token: plainToken,
	}

	//mail
	if err := utils.ResponseJSON(w, http.StatusCreated, userWithToken); err != nil {
		app.internalServerError(w, r, err)
		return
	}
}
