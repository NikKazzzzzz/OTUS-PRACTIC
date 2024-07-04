package unpacksrt

import "testing"

func TestUnpack(t *testing.T) {
	tests := []struct {
		input    string
		expected string
		err      bool
	}{
		{"a4bc2d5e", "aaaabccddddde", false},
		{"abcd", "abcd", false},
		{"45", "", true},
		{"qwe\\4\\5", "qwe45", false},
		{"qwe\\45", "qwe44444", false},
		{"qwe\\\\5", "qwe\\\\\\\\\\", false},
	}

	for _, test := range tests {
		result, err := Unpack(test.input)
		if (err != nil) != test.err {
			t.Errorf("Unpack(%s) error = %v, expected error = %v", test.input, err, test.err)
		}
		if result != test.expected {
			t.Errorf("Unpack(%s) = %s, expected = %s", test.input, result, test.expected)
		}
	}
}
