package files

import (
	"os"
	"regexp"
)

func ExpandWindows(path string) string {
	re := regexp.MustCompile(`\%(?i)(\w*)\%`)
	match := re.FindStringSubmatch(path)
	if match == nil {
		return path
	}
	newValue := os.Getenv(match[1])
	if newValue == "" {
		return path
	}
	return re.ReplaceAllString(path, newValue)
}
