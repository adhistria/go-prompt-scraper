package metadata

import (
	"errors"
	"strings"
)

var ErrNotFound = errors.New("data not found")

func ConvertUrlToPath(url string) string {
	var fileName string
	if strings.Contains(url, "https") {
		fileName = strings.Replace(url, "https://", "", -1)
	} else {
		fileName = strings.Replace(url, "http://", "", -1)
	}
	fileName = strings.Replace(fileName, "/", "_", -1)
	return fileName
}
