package api

import (
	"net/http"

	"github.com/dineshdb/authnz/internal/user"
	"github.com/dineshdb/authnz/internal/utils"
)

func (app *App) GetMyProfile(w http.ResponseWriter, r *http.Request) {
	// Check if user has profile scope
	// TODO: Implement

	utils.OK(w, user.User{})
}
