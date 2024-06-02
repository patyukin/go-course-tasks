package formatter

import (
	"testing"
)

func TestPlainTextFormatter(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"Hello", "Hello"},
		{"world", "world"},
		{"", ""},
		{"12345", "12345"},
		{"!@#$%", "!@#$%"},
	}

	formatter := PlainTextFormatter{}

	for _, test := range tests {
		result := formatter.Format(test.input)
		if result != test.expected {
			t.Errorf("Expected %s, but got %s", test.expected, result)
		}
	}
}

func TestBoldFormatter(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"Hello", "**Hello**"},
		{"world", "**world**"},
		{"", "****"},
		{"12345", "**12345**"},
		{"!@#$%", "**!@#$%**"},
	}

	formatter := BoldFormatter{}

	for _, test := range tests {
		result := formatter.Format(test.input)
		if result != test.expected {
			t.Errorf("Expected %s, but got %s", test.expected, result)
		}
	}
}

func TestCodeFormatter(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"Hello", "`Hello`"},
		{"world", "`world`"},
		{"", "``"},
		{"12345", "`12345`"},
		{"!@#$%", "`!@#$%`"},
	}

	formatter := CodeFormatter{}

	for _, test := range tests {
		result := formatter.Format(test.input)
		if result != test.expected {
			t.Errorf("Expected %s, but got %s", test.expected, result)
		}
	}
}

func TestItalicFormatter(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"Hello", "_Hello_"},
		{"world", "_world_"},
		{"", "__"},
		{"12345", "_12345_"},
		{"!@#$%", "_!@#$%_"},
	}

	formatter := ItalicFormatter{}

	for _, test := range tests {
		result := formatter.Format(test.input)
		if result != test.expected {
			t.Errorf("Expected %s, but got %s", test.expected, result)
		}
	}
}

func TestChainFormatter(t *testing.T) {
	tests := []struct {
		input      string
		expected   string
		formatters []Formatter
	}{
		{"Hello", "Hello", []Formatter{PlainTextFormatter{}}},
		{"world", "**world**", []Formatter{PlainTextFormatter{}, BoldFormatter{}}},
		{"Hello, world!", "`**Hello, world!**`", []Formatter{PlainTextFormatter{}, BoldFormatter{}, CodeFormatter{}}},
		{"12345", "_`12345`_", []Formatter{PlainTextFormatter{}, CodeFormatter{}, ItalicFormatter{}}},
		{"!@#$%", "_`**!@#$%**`_", []Formatter{PlainTextFormatter{}, BoldFormatter{}, CodeFormatter{}, ItalicFormatter{}}},
	}

	for _, test := range tests {
		chainFormatter := ChainFormatter{}
		for _, formatter := range test.formatters {
			chainFormatter.AddFormatter(formatter)
		}
		result := chainFormatter.Format(test.input)
		if result != test.expected {
			t.Errorf("Expected %s, but got %s", test.expected, result)
		}
	}
}
