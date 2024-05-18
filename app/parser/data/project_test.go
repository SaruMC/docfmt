package data

import "testing"

func TestProject_GetTitle(t *testing.T) {
	project := NewProject()
	project.Title = "TestGetTitle"

	if project.GetTitle() != "TestGetTitle" {
		t.Error("GetTitle() should return the title of the project")
	}
}

func TestProject_GetVersion(t *testing.T) {
	project := NewProject()
	project.Version = "TestGetVersion"

	if project.GetVersion() != "TestGetVersion" {
		t.Error("GetVersion() should return the version of the project")
	}
}

func TestProject_GetDescription(t *testing.T) {
	project := NewProject()
	project.Description = "TestGetDescription"

	if project.GetDescription() != "TestGetDescription" {
		t.Error("GetDescription() should return the description of the project")
	}
}
