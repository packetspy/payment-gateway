package server

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/packetspy/go-payment-gateway/internal/service"
	"github.com/packetspy/go-payment-gateway/internal/web/handlers"
)

type Server struct {
	router         *chi.Mux
	server         *http.Server
	accountService *service.AccountService
	port           string
}

func NewServer(accountService *service.AccountService, port string) *Server {
	return &Server{
		router:         chi.NewRouter(),
		accountService: accountService,
		port:           port,
	}
}

func (s *Server) ConfigureRoutes() {
	accountHandler := handlers.NewAccountHandler(s.accountService)

	s.router.Post("/account", accountHandler.Create)
	s.router.Get("/account", accountHandler.Get)
}

func (s *Server) Start() error {
	s.ConfigureRoutes()

	s.server = &http.Server{
		Addr:    ":" + s.port,
		Handler: s.router,
	}

	return s.server.ListenAndServe()
}
