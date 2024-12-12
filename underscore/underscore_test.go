package underscore

import "testing"

func TestCamel(t *testing.T) {
	// Non Anonymous Struct
	// type testCase struct {
	// 	arg  string
	// 	want string
	// }

	// testCases := []testCase{
	// 	{"thisIsACamelCaseString", "this_is_a_camel_case_string"},
	// 	{"a", "a"},
	// }

	// Anonymous Struct
	testCases := []struct {
		arg  string
		want string
	}{
		{"thisIsACamelCaseString", "this_is_a_camel_case_string"},
		{"with a space", "with a space"},
		{"endsWithCapital", "ends_with_capital"},
	}

	for _, tc := range testCases {
		t.Logf("Testing %q", tc.arg)
		got := Camel(tc.arg)
		if got != tc.want {
			t.Errorf("Camel(%q) = %q; want %q", tc.arg, got, tc.want)
		}
	}
}
