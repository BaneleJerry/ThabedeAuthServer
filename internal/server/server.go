package server

import (
	"fmt"
	"net/http"

	"github.com/BaneleJerry/ThabedeAuthServer/config"
	"github.com/BaneleJerry/ThabedeAuthServer/internal/interfaces/database"
)

type Server struct {
	DB         *database.Queries
	HttpServer *http.Server
}

func NewServer(db *database.Queries, serverCfg config.ServerConfig) *Server {

	server := &Server{
		DB: db,
		HttpServer: &http.Server{
			Addr:           fmt.Sprintf(":%s", serverCfg.Port),
			IdleTimeout:    serverCfg.IdleTimeout,
			ReadTimeout:    serverCfg.ReadTimeout,
			WriteTimeout:   serverCfg.WriteTimeout,
			MaxHeaderBytes: serverCfg.MaxHeaderBytes,
		},
	}

	// Initialize routes and set the handler
	router := server.RegisterRoutes()
	server.HttpServer.Handler = router

	return server
}

func (s *Server) Start() {
	fmt.Println("Start Server at Port" + s.HttpServer.Addr)
	if err := s.HttpServer.ListenAndServe(); err != nil {
		fmt.Println("Error: ", err)
	}
}
