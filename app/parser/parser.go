package parser

import (
	"bufio"
	"github.com/misuaaki/docfmt/app/parser/data"
	"os"
	"strings"
)

type Parser struct {
	ProjectInfo data.Project
	FileInfo    data.File
	Functions   []data.Function

	// The current line being parsed
	currentLine string
}

func NewParser() *Parser {
	return &Parser{}
}

func (p *Parser) GetProjectInfo() data.Project {
	return p.ProjectInfo
}

func (p *Parser) GetFileInfo() data.File {
	return p.FileInfo
}

func (p *Parser) GetFunctions() []data.Function {
	return p.Functions
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
	case strings.HasPrefix(line, data.Title):
		p.ProjectInfo.Title = strings.TrimPrefix(line, data.Title)

	case strings.HasPrefix(line, data.Version):
		p.ProjectInfo.Version = strings.TrimPrefix(line, data.Version)

	case strings.HasPrefix(line, data.Description):
		p.ProjectInfo.Description = strings.TrimPrefix(line, data.Description)
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
		params := strings.Split(strings.TrimPrefix(line, data.FunctionParameter), ",")
		for _, param := range params {
			kv := strings.Split(param, " - ")
			if len(kv) == 3 {
				p.Functions[len(p.Functions)-1].Parameters = append(p.Functions[len(p.Functions)-1].Parameters, data.Parameter{Name: kv[0], Value: kv[1], Description: kv[2]})
			}
		}

	case strings.HasPrefix(line, data.FunctionReturn):
		p.createFunctionIfNotExists()
		returns := strings.Split(strings.TrimPrefix(line, data.FunctionReturn), ",")
		for _, ret := range returns {
			ret := strings.Split(ret, " - ")
			p.Functions[len(p.Functions)-1].Returns = append(p.Functions[len(p.Functions)-1].Returns, data.Return{Value: ret[0], Description: ret[1]})
		}

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
