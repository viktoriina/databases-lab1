package app

import (
	"fmt"
	"github.com/pkg/errors"
	"log"
)

func (a *App) updateBlock() error {
	block, err := enterBlock()
	if err != nil {
		return errors.Wrap(err, "failed to enter block")
	}
	if err := a.db.UpdateBlock(block); err != nil {
		return errors.Wrap(err, "failed to update block")
	}
	fmt.Println()
	log.Printf("Updated block with id %d.\n", block.ID)
	return nil
}
