package api

import (
	"encoding/json"
	"net/http"

	"github.com/dineshdb/authnz/internal/user"
	"github.com/dineshdb/authnz/internal/utils"
	"github.com/rs/zerolog/log"
)

type RefreshRequest struct {
	Token string `json:"token"`
}

func (app *App) Refresh(w http.ResponseWriter, r *http.Request) {
	var request RefreshRequest
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		utils.BadRequest(w, "Improper json body")
		return
	}

	id, err := app.JWTValidator.VerifyRefreshToken(request.Token)
	// Invalid refresh token
	if err != nil {
		utils.Unauthorized(w)
		return
	}

	token, err := app.JWTValidator.Generate(user.User{ID: int64(id)})
	if err != nil {
		log.Error().Msg(err.Error())
		utils.Unauthorized(w)
		return
	}
	utils.OK(w, token)
}
