package wzlib_utils

import (
	"strings"
)

// RemovePrefix path from a path with necessary cleaning
func RemovePrefix(pth string, pref string) string {
	if pref != "" && pth != "/" && strings.HasPrefix(pth, pref) {
		pth = "/" + strings.Trim(pth[len(pref):], "/")
	}

	return pth
}
