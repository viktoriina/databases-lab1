package app

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/ethereum/go-ethereum/common"
	"github.com/pkg/errors"
	"github.com/viktoriina/databases-lab1/internal/models"
)

func (a *App) insertTransaction() error {
	transaction, err := enterTransaction()
	if err != nil {
		return errors.Wrap(err, "failed to enter transaction")
	}
	if err := a.db.InsertTransaction(transaction); err != nil {
		return errors.Wrap(err, "failed to insert transaction")
	}
	fmt.Println()
	log.Printf("Inserted transaction with id %d.\n", transaction.ID)
	return nil
}

func enterTransaction() (*models.Transaction, error) {
	reader := bufio.NewReader(os.Stdin)

	id, err := enterTransactionId()
	if err != nil {
		return nil, errors.Wrap(err, "invalid transaction id")
	}

	blockId, err := enterBlockId()
	if err != nil {
		return nil, errors.Wrap(err, "invalid block id")
	}

	fmt.Print("Enter transaction to: ")

	toString, _ := reader.ReadString('\n')
	toString = strings.Replace(toString, "\n", "", replacementLimit)

	if !common.IsHexAddress(toString) {
		return nil, errors.New("invalid transaction to")
	}

	//toString = "0x1234567890123456789012345678901234567890" // Temp for debug purposes
	bytes := common.HexToAddress(toString).Bytes()
	var to [20]byte
	copy(to[:], bytes)

	fmt.Print("Enter transaction value: ")
	valueString, _ := reader.ReadString('\n')
	valueString = strings.Replace(valueString, "\n", "", replacementLimit)
	value, err := strconv.Atoi(valueString)
	if err != nil {
		return nil, errors.Wrap(err, "invalid transaction value")
	}

	return &models.Transaction{
		ID:      id,
		BlockID: blockId,
		To:      to,
		Value:   int64(value),
	}, nil
}
