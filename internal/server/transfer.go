package server

import (
	"errors"
	"net/http"

	"github.com/shopspring/decimal"
)

func (s *server) handleTransfer() http.HandlerFunc {
	type request struct {
		SenderId string `json:"senderId"`
		ReceiverId string `json:"receiverId"`
		Amount string `json:"amount"`
	}
	return func(w http.ResponseWriter, r *http.Request) {
		var request *request
		err := s.decode(w, r, &request)
		if err != nil {
			s.badRequestResponse(w,r, err)
			return
		}
		
		amount, err := decimal.NewFromString(request.Amount)
		if err != nil {
			s.badRequestResponse(w,r, errors.New("invalid amount"))
			return
		}

		err = s.transferService.Transfer(request.SenderId, request.ReceiverId, amount)
		if err != nil {
			s.badRequestResponse(w,r, err)
			return
		}

		s.encode(w, http.StatusOK, envelope{"message": "Transefere done"}, nil)
	}
}
