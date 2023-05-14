package main

import (
	"fmt"
	"os"
	"path/filepath"
	"testing"
)

func Test_run(t *testing.T) {
	tests := []struct {
		name    string
		cfg     config
		wantOut string
	}{
		{
			name:    "DeleteExtensionMatch",
			cfg:     config{ext: ".log", del: true},
			wantOut: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, _ = createTempDir(t, map[string]int{})
		})
	}
}

func createTempDir(t *testing.T, files map[string]int) (dirname string, cleanup func()) {
	t.Helper()
	tempDir, err := os.MkdirTemp("", "walktest")
	if err != nil {
		t.Fatal(err)
	}
	for k, n := range files {
		for j := 1; j <= n; j++ {
			fname := fmt.Sprintf("file%d%s", j, k)
			fpath := filepath.Join(tempDir, fname)
			if err := os.WriteFile(fpath, []byte("dummy"), 0644); err != nil {
				t.Fatal(err)
			}
		}
	}
	return tempDir, func() { os.RemoveAll(tempDir) }
}
