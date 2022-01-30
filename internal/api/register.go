package api

import (
	"encoding/json"
	"net/http"

	"github.com/dineshdb/authnz/internal/user"
	"github.com/dineshdb/authnz/internal/utils"
	"github.com/rs/zerolog/log"
)

type SignupRequest struct {
	User     user.User `json:user`
	Password string    `json:"password"`
}

func (app *App) Signup(w http.ResponseWriter, r *http.Request) {
	var request SignupRequest
	err := json.NewDecoder(r.Body).Decode(&request)

	var newUser user.User = request.User

	if err != nil {
		log.Error().Msg("Bad Request body")
		utils.BadRequest(w, nil)
		return
	}

	// TODO: Implement

	utils.OK(w, newUser)
}
