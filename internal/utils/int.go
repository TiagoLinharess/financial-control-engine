package utils

import "strconv"

func StringToInt64(string string) (int64, error) {
	return strconv.ParseInt(string, 10, 64)
}
