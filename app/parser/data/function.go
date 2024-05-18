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
	Parameters  []Parameter
	Returns     []Return
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
func (f *Function) GetParameters() []Parameter {
	return f.Parameters
}

// GetReturns returns the return values of the function
func (f *Function) GetReturns() []Return {
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

// Parameter represents a parameter of a function
type Parameter struct {
	Name        string
	Value       string
	Description string
}

// GetName returns the parameter of the function
func (p *Parameter) GetName() string {
	return p.Name
}

// GetValue returns the value of the parameter
func (p *Parameter) GetValue() string {
	return p.Value
}

// GetDescription returns the description of the parameter
func (p *Parameter) GetDescription() string {
	return p.Description
}

// Return represents a return value of a function
type Return struct {
	Value       string
	Description string
}

// GetValue returns the value (name) of the return value
func (r *Return) GetValue() string {
	return r.Value
}

// GetDescription returns the description of the return value
func (r *Return) GetDescription() string {
	return r.Description
}
