package data

import (
	"testing"
)

func TestFunction_NewFunction(t *testing.T) {
	f := NewFunction("NewFunction")

	if f == nil {
		t.Error("NewFunction() should create a Function struct")
	}
}

func TestFunction_GetName(t *testing.T) {
	f := NewFunction("GetName")

	if f.GetName() != "GetName" {
		t.Error("GetName() should return the name of the function")
	}
}

func TestFunction_GetDescription(t *testing.T) {
	f := NewFunction("GetDescription")

	f.Description = "description"
	if f.GetDescription() != "description" {
		t.Error("GetDescription should return the name of the function")
	}
}

func TestFunction_GetParameters(t *testing.T) {
	f := NewFunction("GetParameters")

	f.Parameters = map[string]string{"param1": "string", "param2": "int"}
	if len(f.GetParameters()) != 2 {
		t.Error("GetParameters() should return the parameters of the function")
	}
}

func TestFunction_GetReturns(t *testing.T) {
	f := NewFunction("GetReturns")

	f.Returns = []string{"int", "error"}
	if len(f.GetReturns()) != 2 {
		t.Error("GetReturns() should return the return types of the function")
	}
}

func TestFunction_GetComplexity(t *testing.T) {
	f := NewFunction("GetComplexity")

	f.Complexity = "O(n)"
	if f.GetComplexity() != "O(n)" {
		t.Error("GetComplexity() should return the complexity of the function")
	}
}

func TestFunction_GetExample(t *testing.T) {
	f := NewFunction("GetExample")

	f.Example = "Example code"
	if f.GetExample() != "Example code" {
		t.Error("GetExample() should return the example of the function")
	}
}

func TestFunction_GetAuthors(t *testing.T) {
	f := NewFunction("GetAuthors")

	f.Authors = []string{"Author1", "Author2"}
	if len(f.GetAuthors()) != 2 {
		t.Error("GetAuthors() should return the authors of the function")
	}
}
