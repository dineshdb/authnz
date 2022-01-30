package api

import (
	"net/http"

	"github.com/dineshdb/authnz/internal/utils"
)

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (app *App) Login(w http.ResponseWriter, r *http.Request) {
	// TODO: Implement

	utils.OK(w, nil)
}
