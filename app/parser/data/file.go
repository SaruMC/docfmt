package data

// FileDescription represents the general file information
const FileDescription = "@description"

// File represents a file
type File struct {
	Name        string
	Description string
}

// NewFile creates a new File instance
func NewFile(name string) *File {
	return &File{
		Name: name,
	}
}

// GetName returns the name of the file
func (f *File) GetName() string {
	return f.Name
}

// GetDescription returns the description of the file
func (f *File) GetDescription() string {
	return f.Description
}
