package data

const (
	// General project information
	Title       = "@project-title "
	Version     = "@project-version "
	Description = "@project-description "
)

type Project struct {
	Title       string
	Version     string
	Description string
}

func NewProject() *Project {
	return &Project{}
}

func (p *Project) GetTitle() string {
	return p.Title
}

func (p *Project) GetVersion() string {
	return p.Version
}

func (p *Project) GetDescription() string {
	return p.Description
}
