package scanner

import (
	"testing"
)

func TestNewPackage(t *testing.T) {
	p := NewPackage("testPath")

	if p == nil {
		t.Error("NewPackage() should not return nil")
	}

	if p.Path != "testPath" {
		t.Errorf("Expected path to be 'testPath', got %s", p.Path)
	}

	if p.Files == nil {
		t.Error("NewPackage().Files should not be nil")
	}

	if len(p.Files) != 0 {
		t.Error("NewPackage().Files should be empty")
	}
}

func TestGetFiles(t *testing.T) {
	p := NewPackage("testPath")
	p.Files["file1.go"] = nil
	p.Files["file2.go"] = nil

	files := p.GetFiles()

	if len(files) != 2 {
		t.Errorf("Expected 2 files, got %d", len(files))
	}

	if files[0] != "file1.go" && files[1] != "file1.go" {
		t.Error("Expected 'file1.go' to be in the files")
	}

	if files[0] != "file2.go" && files[1] != "file2.go" {
		t.Error("Expected 'file2.go' to be in the files")
	}
}

func TestString(t *testing.T) {
	p := NewPackage("testPath")
	p.Files["file1.go"] = nil
	p.Files["file2.go"] = nil

	expected := "Files: [file1.go file2.go]"
	got := p.String()

	if got != expected {
		t.Errorf("Expected '%s', got '%s'", expected, got)
	}
}
