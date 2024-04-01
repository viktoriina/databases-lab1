package config

import (
	"github.com/pkg/errors"
	"time"
)

func (c *config) BlocksIndexSortPeriod() time.Duration {
	result, err := time.ParseDuration(c.BlocksIndexSortPeriodValue)
	if err != nil {
		panic(errors.Wrap(err, "failed to parse blocks index sort period value"))
	}
	return result
}
