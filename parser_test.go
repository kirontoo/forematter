package forematter

import (
	"bufio"
	"reflect"
	"strings"
	"testing"
)

func TestDetect(t *testing.T) {
	t.Run("detects YAML format", func(t *testing.T) {
		const yamlInput = `---
		hi: hello
		---
		`
		reader := bufio.NewReader(strings.NewReader(yamlInput))
		got, _ := detect(reader)
		if eq := reflect.DeepEqual(got, YamlFormat); !eq {
			t.Fail()
		}
	})

	t.Run("detects TOML format", func(t *testing.T) {
		const tomlInput = `+++
		hi: hello
		+++
		`
		reader := bufio.NewReader(strings.NewReader(tomlInput))
		got, _ := detect(reader)
		if eq := reflect.DeepEqual(got, TomlFormat); !eq {
			t.Fail()
		}
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
		// TODO: test passes when it shouldn't
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

func TestExtract(t *testing.T) {
	const input = `---
		title: my new title
		pubDate: 01-24-2024
		description: this is a simple description
		---
		`

	type Frontmatter struct {
		Title       string `yaml:"title"`
		PubDate     string `yaml:"pubDate"`
		Description string `yaml:"description"`
	}

	got := Frontmatter{}
	expect := Frontmatter{Title: "my new title", PubDate: "01-24-2024", Description: "this is a simple description"}

	reader := bufio.NewReader(strings.NewReader(input))
	extract(reader, YamlFormat, &got)
	if valid := reflect.DeepEqual(got, expect); !valid {
		t.Fail()
	}
}
