package database

import (
	"log"
	"time"

	"github.com/pkg/errors"
	"github.com/viktoriina/databases-lab1/internal/config"
	"github.com/viktoriina/databases-lab1/internal/database/index"
	"github.com/viktoriina/databases-lab1/internal/utils"
)

type offset = int64

type Database struct {
	cfg config.Config

	blocksIndex         *index.Index
	blocksGarbage       []offset
	transactionsIndex   *index.Index
	transactionsGarbage []offset
}

func NewDatabase(cfg config.Config) *Database {
	return &Database{cfg: cfg}
}

func (db *Database) Init() error {
	if err := db.initStorageFiles(); err != nil {
		return err
	}
	if err := db.initIndexes(); err != nil {
		return err
	}
	return nil
}

func (db *Database) initStorageFiles() error {
	if err := utils.EnsureFileExists(db.cfg.BlocksFile()); err != nil {
		return errors.Wrap(err, "failed to ensure blocks file existence")
	}
	if err := utils.EnsureFileExists(db.cfg.TransactionsFile()); err != nil {
		return errors.Wrap(err, "failed to ensure transactions file existence")
	}
	return nil
}

func (db *Database) initIndexes() error {
	var err error
	db.blocksIndex, err = index.NewIndex(db.cfg.BlocksIndexFile())
	if err != nil {
		return errors.Wrap(err, "failed to create blocks index")
	}
	db.transactionsIndex, err = index.NewIndex(db.cfg.TransactionsIndexFile())
	if err != nil {
		return errors.Wrap(err, "failed to create transactions index")
	}
	return nil
}

func (db *Database) Shutdown() error {
	return db.writeIndexes()
}

func (db *Database) writeIndexes() error {
	log.Println("Executing writing to blocks index table..")
	if err := db.blocksIndex.WriteIndexTable(); err != nil {
		return errors.Wrap(err, "failed to write blocks index table")
	}
	log.Println("Writing to blocks index table completed.")

	log.Println("Executing writing to transactions index table..")
	if err := db.transactionsIndex.WriteIndexTable(); err != nil {
		return errors.Wrap(err, "failed to write transactions index table")
	}
	log.Println("Writing to blocks transactions table completed.")

	return nil
}

func (db *Database) RunIndexesSorting() {
	go db.runBlocksIndexSorting()
	go db.runTransactionsIndexSorting()
}

func (db *Database) runBlocksIndexSorting() {
	duration := db.cfg.BlocksIndexSortPeriod()
	for {
		db.blocksIndex.SortIndexTable()
		time.Sleep(duration)
	}
}

func (db *Database) runTransactionsIndexSorting() {
	duration := db.cfg.TransactionsIndexSortPeriod()
	for {
		db.transactionsIndex.SortIndexTable()
		time.Sleep(duration)
	}
}
