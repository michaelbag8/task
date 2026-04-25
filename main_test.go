package main

import (
	"os"
	"strings"
	"testing"
)

// CountLinesInText
func TestCountLinesInText(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  int
	}{
		{"empty string",        "",                    0},
		{"single line",         "Go is powerful",      1},
		{"two lines",           "Go\nPython",          2},
		{"three lines",         "One\nTwo\nThree",     3},
		{"trailing newline",    "Go\nPython\n",        3},
		{"only newline",        "\n",                  2},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := CountLinesInText(tt.input)
			if got != tt.want {
				t.Errorf("CountLinesInText(%q) = %d, want %d", tt.input, got, tt.want)
			}
		})
	}
}

// CountWords
func TestCountWords(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  int
	}{
		{"empty string",      "",                  0},
		{"single word",       "Go",                1},
		{"simple sentence",   "Go is powerful",    3},
		{"extra spaces",      "Hello   there",     2},
		{"newline separated", "One\nTwo\nThree",   3},
		{"tabs as spaces",    "Go\tis\tgreat",     3},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := CountWords(tt.input)
			if got != tt.want {
				t.Errorf("CountWords(%q) = %d, want %d", tt.input, got, tt.want)
			}
		})
	}
}


// ReplaceAllInText
func TestReplaceAllInText(t *testing.T) {
	tests := []struct {
		name    string
		text    string
		oldWord string
		newWord string
		want    string
	}{
		{"simple replace",       "Go is great",         "Go",    "Golang", "Golang is great"},
		{"replace all matches",  "Go Go Go",            "Go",    "Rust",   "Rust Rust Rust"},
		{"no match",             "Go is great",         "Java",  "Rust",   "Go is great"},
		{"empty string",         "",                    "Go",    "Rust",   ""},
		{"replace with empty",   "Go is great",         "great", "",       "Go is "},
		{"old and new same",     "Go is great",         "Go",    "Go",     "Go is great"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ReplaceAllInText(tt.text, tt.oldWord, tt.newWord)
			if got != tt.want {
				t.Errorf("ReplaceAllInText(%q, %q, %q) = %q, want %q",
					tt.text, tt.oldWord, tt.newWord, got, tt.want)
			}
		})
	}
}

// ReplaceWordInFile
func createTempFile(t *testing.T, content string) string {
	t.Helper()
	f, err := os.CreateTemp(t.TempDir(), "testfile-*.txt")
	if err != nil {
		t.Fatalf("failed to create temp file: %v", err)
	}
	if _, err := f.WriteString(content); err != nil {
		t.Fatalf("failed to write to temp file: %v", err)
	}
	f.Close()
	return f.Name()
}

func TestReplaceWordInFile(t *testing.T) {
	tests := []struct {
		name        string
		fileContent string
		oldWord     string
		newWord     string
		want        string
	}{
		{
			name:        "simple replace",
			fileContent: "Go is great",
			oldWord:     "Go",
			newWord:     "Golang",
			want:        "Golang is great",
		},
		{
			name:        "replace all occurrences",
			fileContent: "Go Go Go",
			oldWord:     "Go",
			newWord:     "Rust",
			want:        "Rust Rust Rust",
		},
		{
			name:        "no match",
			fileContent: "Go is great",
			oldWord:     "Java",
			newWord:     "Rust",
			want:        "Go is great",
		},
		{
			name:        "multiline file",
			fileContent: "Go is great\nGo is fast",
			oldWord:     "Go",
			newWord:     "Golang",
			want:        "Golang is great Golang is fast",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			filePath := createTempFile(t, tt.fileContent)
			got := ReplaceWordInFile(filePath, tt.newWord, tt.oldWord)
			if got != tt.want {
				t.Errorf("ReplaceWordInFile() = %q, want %q", got, tt.want)
			}
		})
	}
}

// ValidateASCIIInput

func captureStderr(t *testing.T, fn func()) string {
	t.Helper()
	r, w, err := os.Pipe()
	if err != nil {
		t.Fatalf("os.Pipe: %v", err)
	}
	oldStderr := os.Stderr
	os.Stderr = w

	fn()

	w.Close()
	os.Stderr = oldStderr

	var buf strings.Builder
	tmp := make([]byte, 1024)
	for {
		n, _ := r.Read(tmp)
		if n == 0 {
			break
		}
		buf.Write(tmp[:n])
	}
	r.Close()
	return buf.String()
}

// helper: captures stdout output
func captureStdout(t *testing.T, fn func()) string {
	t.Helper()
	r, w, err := os.Pipe()
	if err != nil {
		t.Fatalf("os.Pipe: %v", err)
	}
	oldStdout := os.Stdout
	os.Stdout = w

	fn()

	w.Close()
	os.Stdout = oldStdout

	var buf strings.Builder
	tmp := make([]byte, 1024)
	for {
		n, _ := r.Read(tmp)
		if n == 0 {
			break
		}
		buf.Write(tmp[:n])
	}
	r.Close()
	return buf.String()
}

func TestValidateASCIIInput(t *testing.T) {
	tests := []struct {
		name       string
		input      string
		wantStdout string
		wantStderr string
	}{
		{
			name:       "valid ascii",
			input:      "Hello!",
			wantStdout: "Valid input\n",
			wantStderr: "",
		},
		{
			name:       "valid with symbols",
			input:      "Go ~ is # great",
			wantStdout: "Valid input\n",
			wantStderr: "",
		},
		{
			name:       "invalid emoji",
			input:      "Hello😊",
			wantStdout: "",
			wantStderr: "Invalid character found: 😊\n",
		},
		{
			name:       "empty string",
			input:      "",
			wantStdout: "Valid input\n",
			wantStderr: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			stdout := captureStdout(t, func() {
				stderr := captureStderr(t, func() {
					ValidateASCIIInput(tt.input)
				})
				if stderr != tt.wantStderr {
					t.Errorf("stderr = %q, want %q", stderr, tt.wantStderr)
				}
			})
			if stdout != tt.wantStdout {
				t.Errorf("stdout = %q, want %q", stdout, tt.wantStdout)
			}
		})
	}
}