package app

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/pkg/errors"
	"github.com/viktoriina/databases-lab1/internal/models"
)

func (a *App) insertBlock() error {
	block, err := enterBlock()
	if err != nil {
		return errors.Wrap(err, "failed to enter block")
	}
	if err := a.db.InsertBlock(block); err != nil {
		return errors.Wrap(err, "failed to insert block")
	}
	fmt.Println()
	log.Printf("Inserted block with id %d.\n", block.ID)
	return nil
}

func enterBlock() (*models.Block, error) {
	reader := bufio.NewReader(os.Stdin)

	id, err := enterBlockId()
	if err != nil {
		return nil, errors.Wrap(err, "invalid block id")
	}

	fmt.Print("Enter block timestamp: ")
	timestampString, _ := reader.ReadString('\n')
	timestampString = strings.Replace(timestampString, "\n", "", replacementLimit)
	timestamp, err := strconv.Atoi(timestampString)
	if err != nil || timestamp < 1 {
		return nil, errors.Wrap(err, "invalid block timestamp")
	}

	fmt.Print("Enter block nonce: ")
	nonceString, _ := reader.ReadString('\n')
	nonceString = strings.Replace(nonceString, "\n", "", replacementLimit)
	nonce, err := strconv.Atoi(nonceString)
	if err != nil {
		return nil, errors.Wrap(err, "invalid block nonce")
	}

	return &models.Block{
		ID:        id,
		Timestamp: int64(timestamp),
		Nonce:     int64(nonce),
	}, nil
}
