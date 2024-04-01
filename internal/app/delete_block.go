package app

import (
	"fmt"
	"github.com/pkg/errors"
	"log"
)

func (a *App) deleteBlock() error {
	id, err := enterBlockId()
	if err != nil {
		return errors.Wrap(err, "invalid block id")
	}
	if err := a.db.DeleteBlock(id); err != nil {
		return errors.Wrap(err, "failed to delete block")
	}
	fmt.Println()
	log.Printf("Deleted block with id %d.\n", id)
	return nil
}
