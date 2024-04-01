package app

import (
	"fmt"

	"github.com/pkg/errors"
	"github.com/viktoriina/databases-lab1/internal/models"
)

func (a *App) dumpTransactionsPerBlock() error {
	blockId, err := enterBlockId()
	if err != nil {
		return errors.Wrap(err, "invalid block id")
	}
	transactions, err := a.db.GetTransactionsPerBlock(blockId, false)
	if err != nil {
		return errors.Wrap(err, "failed to dump transactions")
	}
	for _, tx := range transactions {
		printFullTransaction(&tx)
		fmt.Println()
	}
	return nil
}

func printFullTransaction(tx *models.Transaction) {
	printTransaction(tx)
	fmt.Printf("Transaction logical presence: %t\n", tx.LogicalPresent)
}
