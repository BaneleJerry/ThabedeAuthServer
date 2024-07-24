package api

import (
	"encoding/json"
	"net/http"

	"github.com/BaneleJerry/ThabedeAuthServer/internal/domain/auth"
	"github.com/BaneleJerry/ThabedeAuthServer/pkg/utils"
)

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type RegisterRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

type AuthHandler struct {
	Service auth.AuthService
}

func NewAuthHandler(service auth.AuthService) *AuthHandler {
	return &AuthHandler{Service: service}
}

func (u *AuthHandler) RegisterHandler(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)

	defer r.Body.Close()
	var parameter RegisterRequest

	if err := decoder.Decode(&parameter); err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, "Failed to decode json")
		return
	}

	user, err := u.Service.Register(r.Context(), parameter.Username, parameter.Email, parameter.Password)
	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	utils.RespondWithJSON(w, http.StatusOK, user)

}
func (u *AuthHandler) LoginHandler(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)

	defer r.Body.Close()
	var parameter LoginRequest

	if err := decoder.Decode(&parameter); err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, "Failed to decode json")
		return
	}

	user, err := u.Service.Login(r.Context(), parameter.Email, parameter.Password)
	if err != nil {
		utils.RespondWithError(w, http.StatusUnauthorized, err.Error())
		return
	}

	utils.RespondWithJSON(w, http.StatusOK, user)
}
