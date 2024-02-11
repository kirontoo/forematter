package forematter

import (
	"bufio"
	"reflect"
	"strings"
	"testing"
)

func TestDetect(t *testing.T) {
	t.Run("detects YAML format", func(t *testing.T) {
		const yamlInput = `----
		hi: hello
		----
		`
		reader := bufio.NewReader(strings.NewReader(yamlInput))
		got, _ := detect(reader)
		reflect.DeepEqual(got, YamlFormat)
	})

	t.Run("detects TOML format", func(t *testing.T) {
		const tomlInput = `++++
		hi: hello
		++++
		`
		reader := bufio.NewReader(strings.NewReader(tomlInput))
		got, _ := detect(reader)
		reflect.DeepEqual(got, TomlFormat)
	})

	t.Run("detects invalid delimiter", func(t *testing.T) {
		const input = `;;;
		hi: hello
		;;;`

		reader := bufio.NewReader(strings.NewReader(input))
		got, err := detect(reader)

		if got != nil {
			t.Fail()
		}
		if err == nil {
			t.Fail()
		}
	})

	t.Run("detects invalid format", func(t *testing.T) {
		const input = `++++
		hi: hello
		++++
		`

		reader := bufio.NewReader(strings.NewReader(input))
		got, err := detect(reader)

		if got != nil {
			t.Fail()
		}
		if err == nil {
			t.Fail()
		}
	})
}
