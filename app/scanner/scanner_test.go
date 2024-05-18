package scanner

import (
	"os"
	"path/filepath"
	"testing"
)

func TestNewScanner(t *testing.T) {
	s := NewScanner()

	if s == nil {
		t.Error("NewScanner() should not return nil")
	}

	if s.Packages == nil {
		t.Error("NewScanner().Packages should not be nil")
	}

	if len(s.Packages) != 0 {
		t.Error("NewScanner().Packages should be empty")
	}
}

func TestScan(t *testing.T) {
	s := NewScanner()
	err := s.Scan(".", "go")

	if err != nil {
		t.Errorf("Scan() returned error: %v", err)
	}

	if len(s.Packages) == 0 {
		t.Error("Scan() should populate Packages")
	}

	for path, pkg := range s.Packages {
		if path != pkg.Path {
			t.Errorf("Package path does not match key in Packages map: got %v, want %v", pkg.Path, path)
		}

		if len(pkg.Files) == 0 {
			t.Errorf("Package at path %v has no files", path)
		}

		for filePath, file := range pkg.Files {
			if file == nil {
				t.Errorf("File at path %v is nil", filePath)
			}

			if !filepath.IsAbs(filePath) {
				t.Errorf("File path is not absolute: %v", filePath)
			}

			if _, err := os.Stat(filePath); os.IsNotExist(err) {
				t.Errorf("File at path %v does not exist", filePath)
			}
		}
	}
}
