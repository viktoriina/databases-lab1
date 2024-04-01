package database

import (
	"bytes"
	"encoding/binary"
	"os"

	"github.com/pkg/errors"
	"github.com/viktoriina/databases-lab1/internal/models"
)

var ErrTransactionIsNotPresent = errors.New("transaction is not present")

func (db *Database) GetTransaction(id int64, logical bool) (tx *models.Transaction, errClose error) {
	index, err := db.transactionsIndex.GetIndexById(id)
	if err != nil {
		return nil, err
	}
	if logical && !index.Exists {
		return nil, ErrTransactionIsNotPresent
	}
	file, err := os.OpenFile(db.cfg.TransactionsFile(), os.O_RDONLY, os.ModePerm)
	if err != nil {
		return nil, err
	}
	defer func() {
		if err := file.Close(); err != nil {
			errClose = err
		}
	}()
	tx = new(models.Transaction)
	size := binary.Size(tx)
	txBytes := make([]byte, size)
	if _, err := file.ReadAt(txBytes, index.Offset); err != nil {
		return nil, err
	}
	buf := bytes.NewReader(txBytes)
	if err := binary.Read(buf, binary.BigEndian, tx); err != nil {
		return nil, err
	}
	return tx, nil
}
