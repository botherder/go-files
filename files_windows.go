package files

import (
	"regexp"
)

func ExpandWindows(path string) string {
	re := regexp.MustCompile(`\%(?i)(\w*)\%`)
	match := re.FindStringSubmatch(path)
	newValue := os.Getenv(match[1])
	return re.ReplaceAllString(path, newValue)
}
