package strutil

import (
	"testing"
)

func TestGetURLPathWithoutRootDomain(t *testing.T) {
	tests := []struct {
		rawURL   string
		expected string
	}{
		{"https://example.com/path/to/resource", "resource"},
		{"http://example.com/another/path?something", "path"},
		{"https://example.com/", ""},
		{"invalid-url", "invalid-url"},
		{"invalid-url?blablabla", "invalid-url"},
	}

	for _, test := range tests {
		result := GetLastPartOfURL(test.rawURL)
		if result != test.expected {
			t.Errorf("for URL %s, expected %s, but got %s", test.rawURL, test.expected, result)
		}
	}
}
func TestTransliterate(t *testing.T) {
	tests := []struct {
		input    string
		limit    *int
		expected string
	}{
		{"привет", nil, "privet"},
		{"Привет", nil, "Privet"},
		{"Привет, мир!", nil, "Privet, mir!"},
		{"Привет, мир!", intPtr(10), "Privet, mi"},
		{"", nil, ""},
		{"123", nil, "123"},
		{"Привет123", nil, "Privet123"},
	}

	for _, test := range tests {
		result := Transliterate(test.input, test.limit)
		if result != test.expected {
			t.Errorf("for input %s with limit %v, expected %s, but got %s", test.input, test.limit, test.expected, result)
		}
	}
}

func TestFilterAcceptableChars(t *testing.T) {
	tests := []struct {
		input    string
		limit    *int
		expected string
	}{
		{"hello_world-123", nil, "hello_world-123"},
		{"hello@world!", nil, "helloworld"},
		{"1234567890abcdefghijklmnopqrstuvwxyz", intPtr(10), "1234567890"},
		{"ABCDEFGHIJKLMNOPQRSTUVWXYZ", nil, "ABCDEFGHIJKLMNOPQRSTUVWXYZ"},
		{"!@#$%^&*()_+", nil, "_"},
		{"", nil, ""},
		{"Привет, мир!", nil, ""},
		{"Привет123", nil, "123"},
	}

	for _, test := range tests {
		result := FilterAcceptableChars(test.input, test.limit)
		if result != test.expected {
			t.Errorf("for input %s with limit %v, expected %s, but got %s", test.input, test.limit, test.expected, result)
		}
	}
}

func TestTrimByChars(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{" hello, world. ", "hello, world"},
		{"no,spaces,here,,,", "no,spaces,here"},
		{"trim, these, chars.", "trim, these, chars"},
		{"   alreadyclean   ", "alreadyclean"},
		{" , . ", ""},
		{"", ""},
	}

	for _, test := range tests {
		result := TrimByChars(test.input)
		if result != test.expected {
			t.Errorf("for input %s, expected %s, but got %s", test.input, test.expected, result)
		}
	}
}

func intPtr(i int) *int {
	return &i
}
