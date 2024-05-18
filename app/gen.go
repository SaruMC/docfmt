package app

import (
	"fmt"
	"github.com/misuaaki/godoc/app/parser"
	"github.com/misuaaki/godoc/app/scanner"
	"os"
)

// GenerateTextualDocumentation generates documentation for the given path in a simple text format
// This is used for testing purposes only
func GenerateTextualDocumentation(givenPath string) error {
	s := scanner.NewScanner()
	if err := s.Scan(givenPath, "js"); err != nil {
		return err
	}

	p := parser.NewParser()

	// Create the output file
	file, err := os.Create("output.txt")
	if err != nil {
		return err
	}
	defer file.Close()

	// Iterate over the packages and files
	for _, pkg := range s.GetPackages() {
		for filePath := range pkg.Files {
			// Parse the file
			err := p.Parse(filePath)
			if err != nil {
				return err
			}

			// Write the project info to the output file
			_, err = file.WriteString(fmt.Sprintf("Project: %s\nVersion: %s\nDescription: %s\n\n", p.ProjectInfo.Title, p.ProjectInfo.Version, p.ProjectInfo.Description))
			if err != nil {
				return err
			}

			// Write the file info to the output file
			_, err = file.WriteString(fmt.Sprintf("File: %s\nDescription: %s\n\n", filePath, p.FileInfo.Description))
			if err != nil {
				return err
			}

			// Write the function info to the output file
			for _, function := range p.Functions {
				_, err = file.WriteString(fmt.Sprintf("Function: %s\nDescription: %s\n", function.Name, function.Description))
				if err != nil {
					return err
				}

				_, err = file.WriteString("Parameters: [")
				for _, param := range function.Parameters {
					_, err = file.WriteString(fmt.Sprintf("{Name: %s, Value: %s, Description: %s}, ", param.Name, param.Value, param.Description))
					if err != nil {
						return err
					}
				}
				_, err = file.WriteString("]\n")
				if err != nil {
					return err
				}

				_, err = file.WriteString("Returns: [")
				for _, ret := range function.Returns {
					_, err = file.WriteString(fmt.Sprintf("{type: %s, description: %s}, ", ret.Value, ret.Description))
					if err != nil {
						return err
					}
				}
				_, err = file.WriteString("]\n")
				if err != nil {
					return err
				}

				_, err = file.WriteString(fmt.Sprintf("Complexity: %s\nExample: %s\nAuthor: %s\n\n", function.Complexity, function.Example, function.Authors))
				if err != nil {
					return err
				}
			}
		}
	}

	return nil
}
