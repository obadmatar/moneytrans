package server

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func (s *server) handleAccountsGet() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		accounts, err := s.accountRepository.List()
		if err != nil {
			s.serverErrorResponse(w,r, err)
		}

		s.encode(w, http.StatusOK, envelope{"accounts": accounts}, nil)
	}
}

func (s *server) handleAccountsGetById() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		params := httprouter.ParamsFromContext(r.Context())
		id := params.ByName("id")

		account, err := s.accountRepository.Get(id)
		if err != nil {
			s.notFoundResponse(w,r)
			return
		}

		s.encode(w, http.StatusOK, envelope{"account": account}, nil)
	}
}