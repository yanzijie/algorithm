package string

import (
	"testing"
)

func TestNaivePatternMatching(t *testing.T) {
	tests := []struct {
		name    string
		text    string
		pattern string
		want    int
	}{
		{
			name:    "Test Case 1: Pattern at the beginning",
			text:    "Hello, world!",
			pattern: "Hello",
			want:    0,
		},
		{
			name:    "Test Case 2: Pattern at the end",
			text:    "Hello, world!",
			pattern: "world!",
			want:    7,
		},
		{
			name:    "Test Case 3: Pattern in the middle",
			text:    "Hello, world!",
			pattern: ", wo",
			want:    5,
		},
		{
			name:    "Test Case 4: Pattern not found",
			text:    "Hello, world!",
			pattern: "bye",
			want:    -1,
		},
		{
			name:    "Test Case 5: Empty pattern",
			text:    "Hello, world!",
			pattern: "",
			want:    0,
		},
		{
			name:    "Test Case 6: Empty text",
			text:    "",
			pattern: "hello",
			want:    -1,
		},
		{
			name:    "Test Case 7: Empty text and pattern",
			text:    "",
			pattern: "",
			want:    0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NaivePatternMatching(tt.text, tt.pattern); got != tt.want {
				t.Errorf("NaivePatternMatching() = %v, want %v", got, tt.want)
			}
		})
	}
}
