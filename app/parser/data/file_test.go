package data

import (
	"testing"
)

func TestFile_NewFile(t *testing.T) {
	file := NewFile("NewFile")

	if file == nil {
		t.Error("NewFile() should return a File struct")
	}
}

func TestFile_GetName(t *testing.T) {
	file := NewFile("GetName")

	if file.Name != "GetName" {
		t.Error("GetName() should return the name of the file")
	}
}

func TestFile_GetDescription(t *testing.T) {
	file := NewFile("GetDescription")

	file.Description = "description"
	if file.GetDescription() != "description" {
		t.Error("GetDescription() should return the new description of the file")
	}
}
