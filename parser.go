package forematter

import (
	"bufio"
	"fmt"
	"strings"
)

func Parse(input string) []byte {
	reader := bufio.NewReader(strings.NewReader(input))

	line, err := reader.ReadBytes('\n')
	if err != nil {
		fmt.Print(err)
	}
	fmt.Print(line)

	return line
}
