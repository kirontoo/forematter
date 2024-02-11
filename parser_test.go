package forematter

import (
	"testing"
	"bufio"
	"strings"
	"reflect"
)

func TestParser(t *testing.T) {
	line := Parse(`----
	hi: hello
	----
	`)
	t.Log(line)
}

func TestDetect(t *testing.T) {
	input := `----
	hi: hello
	----
	`

	t.Run("detect format yaml", func(t *testing.T) {
		reader := bufio.NewReader(strings.NewReader(input))
		got, _ := detect(reader)
		reflect.DeepEqual(got, YamlFormat)
	})
}
