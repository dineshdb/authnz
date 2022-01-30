package api

import (
	"encoding/json"
	"net/http"

	"github.com/dineshdb/authnz/internal/user"
	"github.com/dineshdb/authnz/internal/utils"
	"github.com/rs/zerolog/log"
)

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (app *App) Login(w http.ResponseWriter, r *http.Request) {
	var loginRequest LoginRequest
	err := json.NewDecoder(r.Body).Decode(&loginRequest)

	if err != nil {
		log.Debug().Msg(utils.ErrorBadRequestBody)
		utils.BadRequest(w, utils.ErrorBadRequestBody)
		return
	}

	var userRepository *user.Repository = user.NewUserRepository(&app.DB)

	var user *user.User
	user, err = userRepository.GetByEmail(loginRequest.Email)
	if err != nil {
		// Vague response by design: security by obscurity. Add it for critical paths but don't rely on it only.
		log.Debug().Msg("User not found")
		utils.Unauthorized(w)
		return
	}

	matched, err := utils.ComparePasswordAndHash(loginRequest.Password, user.PasswordHash)
	if err != nil {
		log.Debug().Msg("Password comparision failed")
		utils.Unauthorized(w)
		return
	}

	if !matched {
		log.Debug().Msg("Password mismatch")
		utils.Unauthorized(w)
		return
	}

	token, err := app.JWTValidator.Generate(*user)
	if err != nil {
		log.Error().Msg("JWT token generation failed")
		utils.InternalServerError(w, err)
		return
	}

	utils.OK(w, token)
}
