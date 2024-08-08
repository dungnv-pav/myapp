package utils

import (
	"myapp/logger"
	"time"
)

const (
	maskWords = "***"
	maskLen   = 3
)

// ParseTime String型をtime.Timeに変換する
func ParseTime(f string) (*time.Time, error) {
	jst, err := time.LoadLocation("Asia/Tokyo")
	if err != nil {
		logger.Logger.Error().Str("code", "E9000").Err(err).Send()
		return nil, err
	}
	t, err := time.ParseInLocation("2006-01-02", f, jst)
	if err != nil {
		return nil, err
	}
	return &t, nil
}
