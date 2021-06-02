package main

import "testing"

func Test_nameParsingWorksAsExpected(t *testing.T) {
	cases := map[string]string{
		"v3.0.0-alpha2": "3.0.0",
		"v3.0.0-alpha1": "3.0.0",
		"v2.8.1-pl":     "2.8.1",
		"v2.5.0-rc2":    "2.5.0",
	}

	for n, v := range cases {
		if v != parseVersionFromName(n) {
			t.Errorf("Expected %v got %v", v, parseVersionFromName(n))
		}
	}
}
