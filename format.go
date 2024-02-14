package forematter

import (
	"gopkg.in/yaml.v3"
	"github.com/pelletier/go-toml/v2"
)

type UnmarshalFunc func(data []byte, v interface{}) error

type Format struct {
	Delimiter string
	Unmarshal UnmarshalFunc
}

func newFormat(delim string, unmarshalFunc UnmarshalFunc) *Format {
	return &Format{ Delimiter: delim, Unmarshal: unmarshalFunc}
}

var YamlFormat = newFormat( "---", yaml.Unmarshal )
var TomlFormat = newFormat( "+++", toml.Unmarshal )
