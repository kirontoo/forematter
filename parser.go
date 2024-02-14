package forematter

import (
	"bufio"
	"bytes"
	"errors"
	"io"
	"log"
	"strings"
)

func Parse(input string, v interface{}) {
	reader := bufio.NewReader(strings.NewReader(input))

	// Detect
	format, err := detect(reader)
	if err != nil {
		log.Fatal(err)
	}

	extract(reader, format, &v)
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

func extract(reader *bufio.Reader, format *Format, v interface{}) (bool, error) {
	for {
		str, err := reader.ReadBytes('\n')

		line := strings.TrimSpace(string(str))

		if !strings.Contains(line, format.Delimiter) {
			l := bytes.TrimLeft(str, "\t")
			err := format.Unmarshal(l, v)
			if err != nil {
				return false, err
			}
		}

		atEOF := err == io.EOF

		if err != nil && !atEOF {
			return false, err
		}
		if atEOF {
			break
		}
	}

	return true, nil
}
