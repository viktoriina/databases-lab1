package app

import (
	"fmt"
	"github.com/pkg/errors"
	"log"
)

func (a *App) updateTransaction() error {
	tx, err := enterTransaction()
	if err != nil {
		return errors.Wrap(err, "failed to enter transaction")
	}
	if err := a.db.UpdateTransaction(tx); err != nil {
		return errors.Wrap(err, "failed to update transaction")
	}
	fmt.Println()
	log.Printf("Updated transaction with id %d.\n", tx.ID)
	return nil
}
