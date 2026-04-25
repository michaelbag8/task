package main

import (
	"testing"
)

func TestCountWords(t *testing.T) {
	tests := []struct {
		name  string
		word  string
		count int
	}{
		{"simple sentence", "Go is powerful", 3},
		{"sentence", "Hello   there", 2},
		{"newline sentence", "One\nTwo\nThree", 3},
		{"empty sentence", " ", 0},
		{"larg sentence", "Go is powerful language", 4},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := CountWords(test.word)
			if got != test.count {
				 t.Errorf("CountWords(%q) = %d, want %d", test.word, got, test.count)
			}
		})

	}
}
