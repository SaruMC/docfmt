package scanner

import (
	"github.com/misuaaki/godoc/app/scanner/data"
	"go/ast"
	"go/parser"
	"go/token"
	"os"
	"path/filepath"
	"strings"
)

// Scanner represents a scanner
type Scanner struct {
	Packages map[string]*data.Package
}

// NewScanner creates a new Scanner instance
func NewScanner() *Scanner {
	return &Scanner{Packages: make(map[string]*data.Package)}
}

// GetPackages returns the packages
func (s *Scanner) GetPackages() map[string]*data.Package {
	return s.Packages
}

// Scan scans the given path and populates the Packages map
func (s *Scanner) Scan(path, ext string) error {
	if _, ok := s.Packages[path]; ok {
		return nil
	}

	pkg := &data.Package{Path: path, Files: make(map[string]*ast.File)}
	s.Packages[path] = pkg

	return filepath.Walk(path, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if !info.IsDir() {
			if strings.HasSuffix(path, "."+ext) {
				fset := token.NewFileSet()
				parsedFile, err := parser.ParseFile(fset, path, nil, parser.ParseComments)
				if err != nil {
					return err
				}

				pkg.Files[path] = parsedFile
			}
		}

		return nil
	})
}
