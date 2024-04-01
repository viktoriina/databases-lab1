package database

import (
	"bytes"
	"encoding/binary"
	"os"

	"github.com/pkg/errors"
	"github.com/viktoriina/databases-lab1/internal/models"
)

var ErrBlockIsNotPresent = errors.New("block is not present")

func (db *Database) GetBlock(id int64, logical bool) (block *models.Block, errClose error) {
	index, err := db.blocksIndex.GetIndexById(id)
	if err != nil {
		return nil, err
	}
	if logical && !index.Exists {
		return nil, ErrBlockIsNotPresent
	}
	file, err := os.OpenFile(db.cfg.BlocksFile(), os.O_RDONLY, os.ModePerm)
	if err != nil {
		return nil, err
	}
	defer func() {
		if err := file.Close(); err != nil {
			errClose = err
		}
	}()
	block = new(models.Block)
	size := binary.Size(block)
	blockBytes := make([]byte, size)
	if _, err := file.ReadAt(blockBytes, index.Offset); err != nil {
		return nil, err
	}
	buf := bytes.NewReader(blockBytes)
	if err := binary.Read(buf, binary.BigEndian, block); err != nil {
		return nil, err
	}
	return block, nil
}
