package forematter

import (
	"testing"
)

func TestParser(t *testing.T) {
	line := Parse(`----
	hi: hello
	----
	`)
	t.Log(line)
}
