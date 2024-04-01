package app

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/ethereum/go-ethereum/common"
	"github.com/pkg/errors"
	"github.com/viktoriina/databases-lab1/internal/models"
)

func (a *App) getTransaction() error {
	id, err := enterTransactionId()
	if err != nil {
		return errors.Wrap(err, "invalid transaction id")
	}
	tx, err := a.db.GetTransaction(id, true)
	if err != nil {
		return err
	}
	fmt.Println()
	printTransaction(tx)
	return nil
}

func enterTransactionId() (int64, error) {
	fmt.Print("Enter transaction id: ")
	reader := bufio.NewReader(os.Stdin)
	idString, _ := reader.ReadString('\n')
	idString = strings.Replace(idString, "\n", "", replacementLimit)
	id, err := strconv.Atoi(idString)
	const invalidId = -1
	if err != nil || id < 1 {
		return invalidId, err
	}
	return int64(id), nil
}

func printTransaction(tx *models.Transaction) {
	fmt.Println("Transaction id:", tx.ID)
	fmt.Println("Transaction block id:", tx.BlockID)
	fmt.Println("Transaction to:", common.BytesToAddress(tx.To[:]).String())
	fmt.Println("Transaction value:", tx.Value)
}
