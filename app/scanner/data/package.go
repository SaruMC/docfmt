package data

import (
	"fmt"
)

// Package represents a package
type Package struct {
	Path string
	Name string

	Files map[string]string
}

// NewPackage creates a new Package instance
func NewPackage(path string) *Package {
	return &Package{Path: path, Files: make(map[string]string)}
}

// GetFiles returns the files of the package
func (p *Package) GetFiles() []string {
	var files []string

	for filePath := range p.Files {
		files = append(files, filePath)
	}

	return files
}

// String returns the string representation of the package
func (p *Package) String() string {
	return fmt.Sprintf("Files: %v", p.GetFiles())
}
