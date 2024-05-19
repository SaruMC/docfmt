package scanner

import (
	"os"
	"testing"
)

func TestScanner_NewScanner(t *testing.T) {
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

func TestScanner_GetPackages(t *testing.T) {
	s := NewScanner()

	if s.GetPackages() == nil {
		t.Error("GetPackages() should not return nil")
	}

	if len(s.GetPackages()) != 0 {
		t.Error("GetPackages() should return an empty map")
	}

	s.Packages["test"] = nil
	if len(s.GetPackages()) != 1 {
		t.Error("GetPackages() should return a map with one element")
	}

	if s.GetPackages()["test"] != nil {
		t.Error("GetPackages() should return a map with one element")
	}
}

func TestScanner_Scan(t *testing.T) {
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

		for filePath, _ := range pkg.Files {
			if _, err := os.Stat(filePath); os.IsNotExist(err) {
				t.Errorf("File at path %v does not exist", filePath)
			}
		}
	}
}
