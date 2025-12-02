package utils

import (
	"fmt"
	"strings"
	"time"
)

func Time2Str(t time.Time, format string) string {
	if t.IsZero() {
		return ""
	}

	r := strings.NewReplacer(
		"Y", "2006",
		"m", "01",
		"d", "02",
		"H", "15",
		"i", "04",
		"s", "05",
	)
	layout := r.Replace(format)
	str := t.Format(layout)
	if i := strings.Index(str, "t"); i != -1 {
		day := time.Date(t.Year(), t.Month()+1, 1, 0, 0, 0, 0, t.Location()).AddDate(0, 0, -1).Day()
		str = strings.Replace(str, "t", fmt.Sprintf("%02d", day), 1)
	}

	return str
}
