package main

import "testing"

func Test_typeParsingWorksAsExpected(t *testing.T) {
	cases := map[string]string{
		"v3.0.0-alpha2": "alpha",
		"v3.0.0-alpha1": "alpha",
		"v2.8.1-pl":     "stable",
		"v2.5.0-rc2":    "rc",
	}

	for n, v := range cases {
		if v != parseTypeFromName(n) {
			t.Errorf("Expected %v got %v", v, parseTypeFromName(n))
		}
	}
}
