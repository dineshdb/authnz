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
	// NOTICE: Protect against DDoS attacks. Since this endpoint is exposed publicly, it is very easy launch DDoS attacks.
	// Using strong capcha service is recomended.
	// Rate limiting should also be implemented in tandem with captcha to reduce the impact

	if err != nil {
		log.Error().Msg("Bad Request body")
		utils.BadRequest(w, nil)
		return
	}

	passwd, err := utils.HashPasswd(request.Password, &app.ArgonParams)
	if err != nil {
		log.Error().Msg(err.Error())
		utils.InternalServerError(w, nil)
		return
	}

	user, err := user.NewUserRepository(&app.DB).Create(&newUser, passwd)
	if err != nil {
		log.Error().Msg(err.Error())
		utils.BadRequest(w, err)
		return
	}
	utils.OK(w, user)
}
