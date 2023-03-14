package server

import "net/http"

func (s *server) routes() {
	// account routes
	s.router.HandlerFunc(http.MethodGet,"/api/accounts", s.handleAccountsGet())
	s.router.HandlerFunc(http.MethodGet,"/api/accounts/:id", s.handleAccountsGetById())
}