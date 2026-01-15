package time

import (
	"fmt"
	"strings"
	"time"
)

type TimeRange struct {
	Start time.Time
	End   time.Time
}

func NewTimeRange(s string) (*TimeRange, error) {
	const layout = "15:04"
	parts := strings.Split(s, "-")
	if len(parts) != 2 {
		return nil, fmt.Errorf("时间段格式错误")
	}

	loc := time.Now().Location()

	start, err := time.ParseInLocation(layout, strings.TrimSpace(parts[0]), loc)
	if err != nil {
		return nil, err
	}

	end, err := time.ParseInLocation(layout, strings.TrimSpace(parts[1]), loc)
	if err != nil {
		return nil, err
	}

	return &TimeRange{Start: start, End: end}, nil
}

func (tr *TimeRange) Contains(t time.Time) bool {
	start := time.Date(0, 1, 1, tr.Start.Hour(), tr.Start.Minute(), 0, 0, t.Location())
	end := time.Date(0, 1, 1, tr.End.Hour(), tr.End.Minute(), 0, 0, t.Location())
	target := time.Date(0, 1, 1, t.Hour(), t.Minute(), 0, 0, t.Location())

	return !target.Before(start) && !target.After(end)
}
