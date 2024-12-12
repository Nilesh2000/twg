package underscore

import (
	"testing"
)

func TestCamel(t *testing.T) {
	tests := map[string]struct {
		got  string
		want string
	}{
		"simple":            {"thisIsACamelCaseString", "this_is_a_camel_case_string"},
		"spaces":            {"with a space", "with a space1"},
		"ends with capital": {"endsWithA", "ends_with_a"},
		"single capital":    {"A", "_a"},
		"all capitals":      {"ALLCAPS", "_a_l_l_c_a_p_s"},
		"empty string":      {"", ""},
		"no capitals":       {"nocapitals", "nocapitals"},
	}

	for _, tt := range tests {
		t.Run(tt.got, func(t *testing.T) {
			testCamelHelper(t, tt.got, tt.want)
		})
	}
}

func testCamelHelper(t *testing.T, input, expected string) {
	if got := Camel(input); got != expected {
		t.Errorf("Camel() = %v, want %v", got, expected)
	}
}
