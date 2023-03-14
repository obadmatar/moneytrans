package main

import (
	"fmt"
	"os"

	"github.com/obadmatar/moneytrans/internal/account"
	"github.com/obadmatar/moneytrans/internal/server"
	"github.com/obadmatar/moneytrans/internal/transfer"
)

func main() {
	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(1)
	}
}

func run() error {
	accountRepo := account.NewInMemoryRepository("data/accounts.json")

	transferService, err := transfer.NewService(transfer.WithInMemoryAccountRepository(accountRepo))
	if err != nil {
		return err
	}

	srv, err := server.NewServer(server.WithAccountRepository(accountRepo), server.WithTransferService(transferService))
	if err != nil {
		return err
	}

	srv.Run()

	return nil
}
