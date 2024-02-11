package forematter

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"strings"
)

func Parse(input string) {
	reader := bufio.NewReader(strings.NewReader(input))

	// Detect
	_, err := detect(reader)
	if err != nil {
		log.Fatal(err)
	}
}

func detect(reader *bufio.Reader) (*Format, error) {
	firstLine, err := reader.ReadBytes('\n')
	if err != nil {
		fmt.Print(err)
	}

	// Check if YAML or TOML
	if string(firstLine) != YamlFormat.Delimiter || string(firstLine) != TomlFormat.Delimiter {
		return nil, errors.New("invalid format")
	}

	// check for closing delimiter
	// if closing delimiter, then valid input
	scanner := bufio.NewScanner(reader)

	for scanner.Scan() {
		line := scanner.Text()
		if line == YamlFormat.Delimiter {
			return YamlFormat, nil
		}
		if line == TomlFormat.Delimiter {
			return TomlFormat, nil
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return nil, errors.New("invalid format")
}
