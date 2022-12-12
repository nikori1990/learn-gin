package models

import "time"

func UnixToTime(timestamp int) string {
	t := time.Unix(int64(timestamp), 0)
	return t.Format("2006-01-02 15:04:05")
}

func ConcatStr(str1 string, str2 string) string {
	return str1 + "---" + str2
}
