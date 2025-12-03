package utils

import (
	"fmt"
	"strings"
	"time"
)

// 时间转字符串
// 支持 Y 年, m 月, d 日, H 时, i 分, s 秒, t 尾数天
func Time2Str(t time.Time, format string) string {
	if t.IsZero() {
		return ""
	}

	dateTime := t.Format(time.DateTime)
	lastDay := fmt.Sprintf("%2d", time.Date(t.Year(), t.Month()+1, 1, 0, 0, 0, 0, t.Location()).AddDate(0, 0, -1).Day())

	r := strings.NewReplacer(
		"Y", dateTime[:4],
		"m", dateTime[5:7],
		"d", dateTime[8:10],
		"H", dateTime[11:13],
		"i", dateTime[14:16],
		"s", dateTime[17:19],
		"t", lastDay,
	)
	return r.Replace(format)
}

func NowTimeStr(format string) string {
	return Time2Str(time.Now(), format)
}
