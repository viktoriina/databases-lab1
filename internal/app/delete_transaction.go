package app

import (
	"fmt"
	"github.com/pkg/errors"
	"log"
)

func (a *App) deleteTransaction() error {
	id, err := enterTransactionId()
	if err != nil {
		return errors.Wrap(err, "invalid transaction id")
	}
	if err := a.db.DeleteTransaction(id); err != nil {
		return errors.Wrap(err, "failed to delete transaction")
	}
	fmt.Println()
	log.Printf("Deleted transaction with id %d.\n", id)
	return nil
}
