package parser

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

type IniData map[string]map[string]string

type Parser struct {
	data IniData
}

func NewParser() *Parser {
	parser := Parser{}
	parser.data = make(IniData)
	return &parser
}
func (parser *Parser) readFromReader(r io.Reader) error {
	var section string

	sc := bufio.NewScanner(r)

	for sc.Scan() {

		line := strings.Trim(sc.Text(), " \n\t")
		TrimmedLine := strings.TrimSpace(line)
		if len(line) > 2 && line[0] == '[' && line[len(line)-1] == ']' {

			section = strings.TrimSpace(line[1 : len(line)-1])
			parser.data[section] = make(map[string]string)

		} else if strings.Contains(line, "=") {

			keyVal := strings.Split(line, "=")
			key := strings.TrimSpace(keyVal[0])
			val := strings.TrimSpace(keyVal[1])
			parser.data[section][key] = val

		} else if len(TrimmedLine) == 0 {
			continue
		} else if strings.Contains(TrimmedLine, "#") {
			continue
		} else {
			return fmt.Errorf("No sections found")
		}

	}
	return nil
}
func (parser *Parser) ReadFromString(content string) error {
	return parser.readFromReader(strings.NewReader(content))

}

func (parser *Parser) ReadFromFile(path string) error {
	fh, err := os.Open(path)

	if err != nil {
		return fmt.Errorf("Could not open file '%v': %v", path, err)
	}
	return parser.readFromReader(fh)

}

func (parser *Parser) Get(section, key string) (string, error) {
	if parser.data[section][key] == "" {
		return "", fmt.Errorf("Key not found")
	} else {
		return parser.data[section][key], nil

	}
}

func (parser *Parser) GetSection(section string) (map[string]string, error) {
	if parser.data[section] == nil {
		return nil, fmt.Errorf("section not found")
	} else {
		return parser.data[section], nil

	}
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
func (parser *Parser) String() string {
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

func (parser *Parser) WriteToFile(path string) error {
	content := parser.String()

	fh, err := os.Create(path)
	if err != nil {
		return fmt.Errorf("Could not create file '%v': %v", path, err)
	}
	_, err = fh.Write([]byte(content))
	if err != nil {
		return err
	}

	return nil

}
