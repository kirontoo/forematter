package forematter

import (
	"bufio"
	"bytes"
	"errors"
	"log"
	"strings"
)

func Parse(input string, v interface{}) {
	reader := bufio.NewReader(strings.NewReader(input))

	// Detect
	_, err := detect(reader)
	if err != nil {
		log.Fatal(err)
	}
}

// Detect whether or not the format is YAML or TOML
func detect(reader *bufio.Reader) (*Format, error) {
	firstLine, err := reader.ReadBytes('\n')
	if err != nil {
		return nil, err
	}

	line := string(bytes.TrimSpace(firstLine))
	if line == YamlFormat.Delimiter {
		return YamlFormat, nil
	} else if line == TomlFormat.Delimiter {
		return TomlFormat, nil
	}

	return nil, errors.New("invalid format")
}
