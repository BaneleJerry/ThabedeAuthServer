package server

import (
	"net/http"

	"github.com/BaneleJerry/ThabedeAuthServer/internal/domain/auth"
	"github.com/BaneleJerry/ThabedeAuthServer/internal/interfaces/api"
	"github.com/BaneleJerry/ThabedeAuthServer/internal/interfaces/repository"

	"github.com/gorilla/mux"
)

// RegisterRoutes initializes routes for the server
func (s *Server) RegisterRoutes() http.Handler {
	router := mux.NewRouter()

	// Create a subrouter for /auth
	authRouter := router.PathPrefix("/auth").Subrouter()
	authService := api.NewAuthHandler(*auth.NewAuthService(repository.NewUserRepository(s.DB)))

	// Define routes for the subrouter
	authRouter.HandleFunc("/login", authService.LoginHandler).Methods("POST")
	// authRouter.HandleFunc("/logout", authService.LogoutHandler).Methods("POST")
	authRouter.HandleFunc("/register", authService.RegisterHandler).Methods("POST")

	return router
}
