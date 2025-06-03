package utils

import (
	"strconv"
)

func ParseUintParam(param string) (uint, error) {
	val, err := strconv.ParseUint(param, 10, 32)
	return uint(val), err
}
