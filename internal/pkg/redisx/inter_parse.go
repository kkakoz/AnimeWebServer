package redisx

import "strconv"

func InterToInt(v interface{}) (int, bool) {
	if v == nil {
		return 0, false
	}
	vstr, ok := v.(string)
	if !ok {
		return 0, false
	}
	count, err := strconv.Atoi(vstr)
	if err != nil {
		return 0, false
	}
	return count, true
}

func InterToIntF(v interface{}) int {
	if v == nil {
		return 0
	}
	vstr, ok := v.(string)
	if !ok {
		return 0
	}
	count, err := strconv.Atoi(vstr)
	if err != nil {
		return 0
	}
	return count
}
