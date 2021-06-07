package redisx

import "strconv"

func ParseIntFromMap(data map[string]string, key string) int {
	val, ok := data[key]
	if !ok {
		return 0
	}
	i, err := strconv.Atoi(val)
	if err != nil {
		return 0
	}
	return i
}