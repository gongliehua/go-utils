package utils

import (
	"fmt"
	"strings"
	"time"
)

var (
	timeReplacer = map[string]string{
		"Y": "2006",
		"m": "01",
		"d": "02",
		"H": "15",
		"i": "04",
		"s": "05",
	}
)

// 时间转字符串
// 支持 Y 年, m 月, d 日, H 时, i 分, s 秒, t 尾数天
func Time2Str(t time.Time, format string) string {
	if t.IsZero() {
		return ""
	}

	/**
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
	*/

	// fix "Y-m-01"
	str := format
	for k, v := range timeReplacer {
		if strings.Contains(format, k) {
			str = strings.ReplaceAll(str, k, t.Format(v))
		}
	}
	if strings.Contains(str, "t") {
		day := time.Date(t.Year(), t.Month()+1, 1, 0, 0, 0, 0, t.Location()).AddDate(0, 0, -1).Day()
		str = strings.ReplaceAll(str, "t", fmt.Sprintf("%02d", day))
	}

	return str
}

func NowTimeStr(format string) string {
	return Time2Str(time.Now(), format)
}
