package cabdl

import (
	"strings"
	"strconv"
	"github.com/pkg/errors"
	"fmt"
)

type Period struct {
	Year int64
	Month int64
}

func (p Period) String() string {
	return fmt.Sprintf("%d-%02d", p.Year, p.Month)
}

func AsPeriod(s string) (*Period, error) {
	parts := strings.Split(s, "-")

	if len(parts) != 2 {
		return nil, errors.New("period must contain two parts delimited by '-', e.g. 2009-01")
	}

	year, err := strconv.ParseInt(parts[0], 10, 64)
	if err != nil {
		return nil, errors.New("year is not a valid integer")
	}

	month, err := strconv.ParseInt(parts[1], 10, 0)
	if err != nil {
		return nil, errors.New("month is not a valid integer")
	}

	return &Period{
		Year: year,
		Month: month,
	}, nil
}

func ForEachPeriod(start Period, end Period, context *Context, callable func(*Context, Period)) {
	for i := start.Year; i <= end.Year; i++ {
		var endMonth int64

		if i == end.Year {
			endMonth = end.Month
		} else {
			endMonth = 12
		}

		for j := start.Month; j <= endMonth; j++ {
			period := Period{i, j}

			//log.Printf("Entering handler for period %s", period)
			callable(context, period)
			//log.Printf("Leaving handler for period %s", period)
		}
	}
}