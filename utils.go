package main

import (
	"net/url"
)

func maybeURLDecode(s string) string {
	decoded, err := url.QueryUnescape(s)
	if err != nil {
		return s
	}

	return decoded
}
