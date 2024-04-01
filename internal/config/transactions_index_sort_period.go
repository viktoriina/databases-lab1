package config

import (
	"github.com/pkg/errors"
	"time"
)

func (c *config) TransactionsIndexSortPeriod() time.Duration {
	result, err := time.ParseDuration(c.TransactionsIndexSortPeriodValue)
	if err != nil {
		panic(errors.Wrap(err, "failed to parse transactions index sort period value"))
	}
	return result
}
