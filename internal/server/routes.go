package server

import "net/http"

func (s *server) routes() {
	// account routes
	s.router.HandlerFunc(http.MethodGet,"/api/accounts", s.handleAccountsGet())
	s.router.HandlerFunc(http.MethodGet,"/api/accounts/:id", s.handleAccountsGetById())

	// transfer routes
	s.router.HandlerFunc(http.MethodPost, "/api/transfer", s.handleTransfer())
}