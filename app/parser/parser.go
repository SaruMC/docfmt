package parser

import (
	"bufio"
	"github.com/misuaaki/godoc/app/parser/data"
	"os"
	"strings"
)

const (
	// General project information
	title       = "@project-title "
	version     = "@project-version "
	description = "@project-description "
)

type Project struct {
	Title       string
	Version     string
	Description string
}

type Parser struct {
	ProjectInfo Project
	FileInfo    data.File
	Functions   []data.Function

	// The current line being parsed
	currentLine string
}

func NewParser() *Parser {
	return &Parser{}
}

func (p *Parser) Parse(filePath string) error {
	file, err := os.Open(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		if strings.HasPrefix(scanner.Text(), "// ") {
			p.currentLine = scanner.Text()[3:]

		} else if strings.HasPrefix(scanner.Text(), data.FunctionName) {
			p.currentLine = scanner.Text()

		} else {
			continue
		}

		p.parseLine()
	}

	if err := scanner.Err(); err != nil {
		return err
	}

	return nil
}

func (p *Parser) createFunctionIfNotExists() {
	if len(p.Functions) == 0 || p.Functions[len(p.Functions)-1].Name != "" {
		p.Functions = append(p.Functions, *data.NewFunction(""))
	}
}

func (p *Parser) parseLine() {
	line := strings.TrimSpace(p.currentLine)

	// Parse project information
	switch {
	case strings.HasPrefix(line, title):
		p.ProjectInfo.Title = strings.TrimPrefix(line, title)

	case strings.HasPrefix(line, version):
		p.ProjectInfo.Version = strings.TrimPrefix(line, version)

	case strings.HasPrefix(line, description):
		p.ProjectInfo.Description = strings.TrimPrefix(line, description)
	}

	// Parse file information
	switch {
	case strings.HasPrefix(line, data.FileDescription):
		p.FileInfo.Description = strings.TrimPrefix(line, data.FileDescription)
	}

	// Parse function information
	switch {
	case strings.HasPrefix(line, data.FunctionName):
		p.createFunctionIfNotExists()
		p.Functions[len(p.Functions)-1].Name = strings.Split(strings.TrimPrefix(line, data.FunctionName), "(")[0]

	case strings.HasPrefix(line, data.FunctionDescription):
		p.createFunctionIfNotExists()
		p.Functions[len(p.Functions)-1].Description = strings.TrimPrefix(line, data.FunctionDescription)

	case strings.HasPrefix(line, data.FunctionParameter):
		p.createFunctionIfNotExists()
		p.Functions[len(p.Functions)-1].Parameters = make(map[string]string)
		params := strings.Split(strings.TrimPrefix(line, data.FunctionParameter), ",")
		for _, param := range params {
			kv := strings.Split(param, "=")
			if len(kv) == 2 {
				p.Functions[len(p.Functions)-1].Parameters[kv[0]] = kv[1]
			}
		}

	case strings.HasPrefix(line, data.FunctionReturn):
		p.createFunctionIfNotExists()
		p.Functions[len(p.Functions)-1].Returns = strings.Split(strings.TrimPrefix(line, data.FunctionReturn), ",")

	case strings.HasPrefix(line, data.FunctionComplexity):
		p.createFunctionIfNotExists()
		p.Functions[len(p.Functions)-1].Complexity = strings.TrimPrefix(line, data.FunctionComplexity)

	case strings.HasPrefix(line, data.FunctionExample):
		p.createFunctionIfNotExists()
		p.Functions[len(p.Functions)-1].Example = strings.TrimPrefix(line, data.FunctionExample)

	case strings.HasPrefix(line, data.FunctionAuthor):
		p.createFunctionIfNotExists()
		p.Functions[len(p.Functions)-1].Authors = strings.Split(strings.TrimPrefix(line, data.FunctionAuthor), ",")
	}
}
