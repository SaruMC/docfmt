package data

// Function information
const (
	FunctionName        = "function "
	FunctionDescription = "@description "
	FunctionParameter   = "@param "
	FunctionReturn      = "@return "
	FunctionComplexity  = "@complexity "
	FunctionExample     = "@example "
	FunctionAuthor      = "@author "
)

// Function represents a function
type Function struct {
	Name        string
	Description string
	Parameters  map[string]string
	Returns     []string
	Complexity  string
	Example     string
	Authors     []string
}

// NewFunction creates a new Function instance
func NewFunction(name string) *Function {
	return &Function{
		Name: name,
	}
}

// GetName returns the name of the function
func (f *Function) GetName() string {
	return f.Name
}

// GetDescription returns the description of the function
func (f *Function) GetDescription() string {
	return f.Description
}

// GetParameters returns the parameters of the function
func (f *Function) GetParameters() map[string]string {
	return f.Parameters
}

// GetReturns returns the return values of the function
func (f *Function) GetReturns() []string {
	return f.Returns
}

// GetComplexity returns the complexity of the function
func (f *Function) GetComplexity() string {
	return f.Complexity
}

// GetExample returns the example of the function
func (f *Function) GetExample() string {
	return f.Example
}

// GetAuthors returns the authors of the function
func (f *Function) GetAuthors() []string {
	return f.Authors
}
