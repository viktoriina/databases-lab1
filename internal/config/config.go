package config

import (
	"os"
	"time"

	"github.com/pkg/errors"
	"gopkg.in/yaml.v2"
)

type Config interface {
	BlocksFile() string
	BlocksIndexFile() string
	BlocksIndexSortPeriod() time.Duration

	TransactionsFile() string
	TransactionsIndexFile() string
	TransactionsIndexSortPeriod() time.Duration
}

type config struct {
	BlocksFileValue            string `yaml:"blocks_file"`
	BlocksIndexFileValue       string `yaml:"blocks_index_file"`
	BlocksIndexSortPeriodValue string `yaml:"blocks_index_sort_period"`

	TransactionsFileValue            string `yaml:"transactions_file"`
	TransactionsIndexFileValue       string `yaml:"transactions_index_file"`
	TransactionsIndexSortPeriodValue string `yaml:"transactions_index_sort_period"`
}

func NewConfig() (Config, error) {
	const filenameEnv = "config.yaml"
	//filename := os.Getenv(filenameEnv)
	data, err := os.ReadFile(filenameEnv)
	if err != nil {
		return nil, err
	}
	cfg := new(config)
	if err = yaml.Unmarshal(data, cfg); err != nil {
		return cfg, errors.Wrap(err, "failed to init config")
	}
	return cfg, nil
}
