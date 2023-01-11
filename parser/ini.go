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

		line := strings.TrimSpace(sc.Text())
		trimmedline := strings.TrimSpace(line)

		if len(trimmedline) == 0 || strings.HasPrefix(trimmedline, "#") {
			continue
		}
		if len(line) > 2 && line[0] == '[' {
			if line[len(line)-1] == ']' {
				section = strings.TrimSpace(line[1 : len(line)-1])
				parser.data[section] = make(map[string]string)

			} else {
				return fmt.Errorf("invalid section")
			}

		}
		if strings.Contains(line, "=") {

			keyVal := strings.Split(line, "=")
			key := strings.TrimSpace(keyVal[0])
			val := strings.TrimSpace(keyVal[1])
			if key == "" {
				return fmt.Errorf("key is missing")
			}
			if val == "" {
				return fmt.Errorf("value is missing")
			}

			parser.data[section][key] = val
			continue

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
	defer fh.Close()

	return parser.readFromReader(fh)

}

func (parser *Parser) Get(section, key string) (string, error) {
	return parser.data[section][key], nil

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

func (parser *Parser) WriteToFile() error {
	content := parser.String()
	fh, err := os.CreateTemp("", "tmpfile")
	if err != nil {
		return fmt.Errorf("Could not create file : %v", err)
	}
	defer fh.Close()

	_, err = fh.Write([]byte(content))

	return err

}
