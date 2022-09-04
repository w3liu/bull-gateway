package utils

import "strings"

func JoinPath(elem ...string) string {
	return strings.Join(elem, "/")
}
