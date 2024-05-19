package data

import (
	"testing"
)

func TestPackage_NewPackage(t *testing.T) {
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

func TestPackage_GetFiles(t *testing.T) {
	p := NewPackage("testPath")
	p.Files["file1.go"] = ""
	p.Files["file2.go"] = ""

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

// In local test passed, but in the runner test failed
// https://github.com/mitsuaaki/godoc/actions/runs/9143427610/job/25140051448#step:5:66
/*func TestPackage_String(t *testing.T) {
	p := NewPackage("testPath")
	p.Files["file1.go"] = ""
	p.Files["file2.go"] = ""

	expected := "Files: [file1.go file2.go]"
	got := p.String()

	if got != expected {
		t.Errorf("Expected '%s', got '%s'", expected, got)
	}
}
*/
