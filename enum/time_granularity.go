package enum

import (
	"errors"
)

var ErrUnknownTimeGranularity = errors.New("Unknown timeGranularity value")

type TimeGranularity int

const (
	TG_DAYS TimeGranularity = iota
	TG_HOURS
	TG_MINUTES
)

var timeGranularityItems = [...]string{
	"DAYS",
	"HOURS",
	"MINUTES",
}

var timeGranularityMap = map[string]TimeGranularity{
	timeGranularityItems[TG_DAYS]:    TG_DAYS,
	timeGranularityItems[TG_HOURS]:   TG_HOURS,
	timeGranularityItems[TG_MINUTES]: TG_MINUTES,
}

func (code TimeGranularity) String() string {
	return timeGranularityItems[code]
}

func (code *TimeGranularity) UnmarshalJSON(buf []byte) error {
	var err error
	val, ok := timeGranularityMap[string(buf)]
	if ok {
		*code = val
	} else {
		err = ErrUnknownTimeGranularity
	}
	return err

}
