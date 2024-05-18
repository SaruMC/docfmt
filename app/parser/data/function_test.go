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

	f.Parameters = []Parameter{
		{Name: "param1", Value: "int", Description: "A basic integer parameter"},
		{Name: "param2", Value: "string", Description: "A basic string parameter"},
	}

	if len(f.GetParameters()) != 2 {
		t.Error("GetParameters() should return the parameters of the function")
	}

	if f.GetParameters()[0].GetName() != "param1" {
		t.Error("GetParameters() should return the parameters of the function")
	}

	if f.GetParameters()[1].GetName() != "param2" {
		t.Error("GetParameters() should return the parameters of the function")
	}

	if f.GetParameters()[0].GetValue() != "int" {
		t.Error("GetParameters() should return the parameters of the function")
	}

	if f.GetParameters()[1].GetValue() != "string" {
		t.Error("GetParameters() should return the parameters of the function")
	}

	if f.GetParameters()[0].GetDescription() != "A basic integer parameter" {
		t.Error("GetParameters() should return the parameters of the function")
	}

	if f.GetParameters()[1].GetDescription() != "A basic string parameter" {
		t.Error("GetParameters() should return the parameters of the function")
	}
}

func TestFunction_GetReturns(t *testing.T) {
	f := NewFunction("GetReturns")

	f.Returns = []Return{
		{Value: "int", Description: "A basic integer return value"},
		{Value: "error", Description: "An error return value"},
	}

	if len(f.GetReturns()) != 2 {
		t.Error("GetReturns() should return the return types of the function")
	}

	if f.GetReturns()[0].GetValue() != "int" {
		t.Error("GetReturns() should return the return types of the function")
	}

	if f.GetReturns()[1].GetValue() != "error" {
		t.Error("GetReturns() should return the return types of the function")
	}

	if f.GetReturns()[0].GetDescription() != "A basic integer return value" {
		t.Error("GetReturns() should return the return types of the function")
	}

	if f.GetReturns()[1].GetDescription() != "An error return value" {
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
