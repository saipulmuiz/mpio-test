package config

import (
	"fmt"
	"time"

	"github.com/saipulmuiz/mpio-test/pkg/serror"
	"github.com/saipulmuiz/mpio-test/pkg/utils/utstring"
)

func (cfg *Config) InitTimezone() serror.SError {
	loc := utstring.Env("APP_TIMEZONE", "Asia/Jakarta")
	local, err := time.LoadLocation(loc)
	if err != nil {
		return serror.NewFromErrorc(err, fmt.Sprintf("failed load location %s", loc))
	}
	time.Local = local

	return nil
}
