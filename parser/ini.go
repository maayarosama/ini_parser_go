package parser

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Data map[string]map[string]string

type Parser struct {
	data Data
}

func NewParser() *Parser {
	parser := Parser{}
	parser.data = make(Data)
	return &parser
}

func (parser *Parser) ReadFromString(content string) {
	lines := strings.Split(content, "\n")
	var section string

	for _, line := range lines {
		line = strings.TrimSpace(line)
		if len(line) > 2 && line[0] == '[' && line[len(line)-1] == ']' {

			section = strings.TrimSpace(line[1 : len(line)-1])
			parser.data[section] = map[string]string{}

		} else if strings.Contains(line, "=") {
			keyVal := strings.Split(line, " = ")
			key := keyVal[0]
			val := keyVal[1]
			parser.data[section][key] = val

		}
	}
}

func (parser *Parser) ReadFile(path string) error {
	fh, err := os.Open(path)
	var section string

	if err != nil {
		return fmt.Errorf("Could not open file '%v': %v", path, err)
	}

	sc := bufio.NewScanner(fh)

	for sc.Scan() {

		line := strings.Trim(sc.Text(), " \n\t")

		if len(line) > 2 && line[0] == '[' && line[len(line)-1] == ']' {
			section = strings.TrimSpace(line[1 : len(line)-1])
			parser.data[section] = map[string]string{}
		} else if strings.Contains(line, "=") {
			keyVal := strings.Split(line, " = ")
			key := strings.TrimSpace(keyVal[0])
			val := strings.TrimSpace(keyVal[1])
			parser.data[section][key] = val
		}

	}

	return nil
}
func (parser *Parser) Get(section, key string) string {
	return parser.data[section][key]
}

func (parser *Parser) GetSection(section string) map[string]string {
	return parser.data[section]
}
func (parser *Parser) GetSections() []string {
	var sections []string
	for k := range parser.data {
		sections = append(sections, k)
	}
	return sections
}

func (parser *Parser) GetSectionKeys(section string) []string {
	var keys []string
	for k := range parser.data[section] {
		keys = append(keys, k)
	}
	return keys
}
func (parser *Parser) ToString() string {
	var content string

	for sections, keyVal := range parser.data {
		content += fmt.Sprintf("[%s]\n", sections)
		for key, value := range keyVal {
			content += fmt.Sprintf("%s = %s\n", key, value)
		}
		content += "\n"
	}
	return content
}

func (parser *Parser) WriteToFile(path string) (bool, error) {
	content := parser.ToString()

	fh, err := os.Create(path)
	if err != nil {
		return false, fmt.Errorf("Could not create file '%v': %v", path, err)
	}
	_, w_err := fh.Write([]byte(content))
	if w_err != nil {
		return false, w_err
	}

	return true, nil

}
