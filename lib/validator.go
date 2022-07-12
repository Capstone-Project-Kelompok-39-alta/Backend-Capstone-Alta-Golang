package lib

import (
	"strings"
)

func CheckExtensionFile(extension string) bool {
	allowed := []string{
		".csv", ".xls", ".xlsx",
	}
	extension = strings.ToLower(extension)
	for _, v := range allowed {
		if v == extension {
			return true
		}
	}
	return false
}
