package utils

import (
	"testing"
)

func TestNowTimeStr(t *testing.T) {
	t.Log(NowTimeStr("Y-m-d H:i:s"))
	t.Log(NowTimeStr("Y-m-d"))
	t.Log(NowTimeStr("H:i:s"))
	t.Log(NowTimeStr("Y-m-01"))
	t.Log(NowTimeStr("Y-m-t"))
}
