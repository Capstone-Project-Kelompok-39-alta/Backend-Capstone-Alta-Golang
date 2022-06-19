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

func CheckExtensionEmail(extension string) bool {
	allowed := []string{
		"@gmail.com", "@yahoo.com", "@outlook.com", "@yahoo.co.id",
	}
	extension = strings.ToLower(extension)
	for _, v := range allowed {
		if v == extension {
			return true
		}
	}
	return false
}
