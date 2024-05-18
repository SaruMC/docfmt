package parser

import (
	"testing"
)

func TestParser_NewParser(t *testing.T) {
	p := NewParser()

	if p == nil {
		t.Error("Parser is nil")
	}
}

func TestParser_GetFunctions(t *testing.T) {
	p := NewParser()

	if p.GetFunctions() == nil {
		t.Error("Functions is nil")
	}
}
