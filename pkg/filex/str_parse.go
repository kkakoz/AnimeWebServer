package filex

import "strings"

func ParseFileType(filename string) string {
	split := strings.Split(filename, ".")
	length := len(split)
	if length > 1 {
		return split[length - 1]
	}
	return ""
}

var fileTypes = []string{"jpg", "png", "gif"}

func IsImageFile(fileType string) bool {
	lower := strings.ToLower(fileType)
	for _, ftype := range fileTypes {
		if ftype == lower {
			return true
		}
	}
	return false
}