package wzlib_utils

import (
	"net/http"
	"os"
	"strings"
)

// FileExists checks if file... well, exists.
func FileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}

// FileContentType detects the content type via http.
// If unable to detect a proper one, "application/octet-stream" is returned.
func FileContentType(out *os.File) (string, error) {
	buff := make([]byte, 0x200)
	_, err := out.Read(buff)
	if err != nil {
		return "", err
	}

	return strings.Split(http.DetectContentType(buff), ";")[0], nil
}

// FileContentTypeByPath is a convenience alias to FileContentType to handle paths in strings.
func FileContentTypeByPath(pth string) (string, error) {
	fp, err := os.Open(pth)
	if err != nil {
		return "", err
	}
	defer fp.Close()

	return FileContentType(fp)
}
