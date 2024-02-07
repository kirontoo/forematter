# Forematter

## Why
A library to parse frontmatter from input, ignoring the content and spit it out in another format for use.

input: string
output: key-value object (struct)

intended usage:
```go
package main

import "github.com/kirontoo/forematter"

func main() {
    data, err := forematter.Parse(input)
}
```