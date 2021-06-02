package main

import "strings"

func parseVersionFromName(n string) string {
	// get rid of prepended "v"
	if n[:1] == "v" {
		n = n[1:]
	}

	// split on "-"
	parts := strings.Split(n, "-")
	return parts[0]
}
