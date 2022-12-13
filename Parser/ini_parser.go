package main

import (
	"bufio"
	"fmt"
	"os"
)

type Data map[string]map[string]string

type Parser struct {
	data Data
}

func NewParser() Parser {
	parser := Parser{}
	parser.data = make(Data)
	return parser
}
func ReadFile(path string) error {
	fh, err := os.Open(path)

	if err != nil {
		return fmt.Errorf("Could not open file '%v': %v", path, err)
	}
	// sectionHead := regexp.MustCompile(`^\[([^]]*)\]\s*$`)
	// keyValue := regexp.MustCompile(`^(\w*)\s*=\s*(.*?)\s*$`)
	sc := bufio.NewScanner(fh)
	for sc.Scan() {
		_ = sc.Text() // GET the line string
	}
	// file, err := os.ReadFile(path)

	// if err != nil {
	// 	return err
	// }

	// d := string(file)
	// line := strings.Trim(d, " \n\t")
	// fmt.Println(line, "GJG")

	// lines := strings.Split(d, "\n")
	// fmt.Println(lines, "GJG")

	// fmt.Print(string(file))

	return nil
}

func main() {
	// mydir, err := os.Getwd()
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// fmt.Println(mydir)
	ReadFile("../sample.ini")
}
