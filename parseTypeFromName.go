package main

import "strings"

func parseTypeFromName(n string) string {
	parts := strings.Split(n, "-")
	if len(parts) != 2 {
		return ""
	}

	if parts[1] == "pl" {
		return "stable"
	}

	if parts[1][0:2] == "rc" {
		return "rc"
	}

	if parts[1][0:5] == "alpha" {
		return "alpha"
	}
	return ""
}
