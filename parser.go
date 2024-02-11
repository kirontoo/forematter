package forematter

import (
	"bufio"
	"fmt"
	"strings"
	"errors"
)

func Parse(input string) []byte {
	reader := bufio.NewReader(strings.NewReader(input))

	line, err := reader.ReadBytes('\n')
	if err != nil {
		fmt.Print(err)
	}
	fmt.Print(string(line))

	// Detect
	_, err = detect(reader)
	if err != nil {
		fmt.Print(err)
	}

	return line
}

func detect(reader *bufio.Reader) (*Format, error) {
	line, err := reader.ReadBytes('\n')
	if err != nil {
		fmt.Print(err)
	}

	// Check if YAML or TOML
	if string(line) == YamlFormat.Delimiter {
		return YamlFormat, nil
	}

	if string(line) == TomlFormat.Delimiter {
		return TomlFormat, nil
	}

	// check for closing delimiter
	// if closing delimiter, then valid input

	
	return nil, errors.New("invalid format")
}

