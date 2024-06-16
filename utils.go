package main

import (
	"net/url"
	"strings"
)

func maybeURLDecode(s string) string {
	s = strings.TrimSpace(s)
	decoded, err := url.QueryUnescape(s)
	if err != nil {
		return s
	}

	return decoded
}
