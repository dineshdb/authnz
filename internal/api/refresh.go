package api

import (
	"encoding/json"
	"net/http"

	"github.com/dineshdb/authnz/internal/utils"
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

	// TODO: Implement

	utils.OK(w, "token")
}
