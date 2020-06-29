package wzlib_utils

import "strings"

// Ba65Str -- Byte Array of 65 to String extracts a C string from null-terminated array
func Ba65Str(data [65]byte) string {
	var buf [65]byte
	for i, b := range data {
		buf[i] = byte(b)
	}
	str := string(buf[:])
	if i := strings.Index(str, "\x00"); i != -1 {
		str = str[:i]
	}
	return str
}
