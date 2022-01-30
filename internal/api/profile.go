package api

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/dineshdb/authnz/internal/auth"
	"github.com/dineshdb/authnz/internal/user"
	"github.com/dineshdb/authnz/internal/utils"
	"github.com/rs/zerolog/log"
)

func (app *App) GetMyProfile(w http.ResponseWriter, r *http.Request) {
	scope := r.Context().Value(auth.Scope).(string)
	// Check if user has profile scope
	if !strings.Contains(scope, "profile") {
		log.Debug().Msg("user needs to be granted `profile` scope")
		utils.Unauthorized(w)
		return
	}

	userId := r.Context().Value(auth.Subject).(int)
	userProfile, err := user.NewUserRepository(&app.DB).GetByID(int64(userId))
	if err != nil {
		log.Debug().Msg(fmt.Sprintf("%v", err))
		utils.NotFound(w, nil)
		return
	}

	utils.OK(w, userProfile)
}
