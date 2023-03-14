package server

import (
	"log"
	
	"github.com/julienschmidt/httprouter"
	"github.com/obadmatar/moneytrans/internal/account"
	"github.com/obadmatar/moneytrans/internal/transfer"
)

type server struct {
	logger *log.Logger
	router *httprouter.Router
	transferService *transfer.Service
	accountRepository account.Repositoy
}
