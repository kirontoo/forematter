package forematter

import (
	"gopkg.in/yaml.v3"
	"github.com/pelletier/go-toml/v2"
)

type UnMarshalFunc func(data []byte, v interface{}) error

type Format struct {
	Delimiter string
	UnMarshal UnMarshalFunc
}

func newFormat(delim string, unMarshalFunc UnMarshalFunc) *Format {
	return &Format{ Delimiter: delim, UnMarshal: unMarshalFunc}
}

var YamlFormat = newFormat( "---", yaml.Unmarshal )
var TomlFormat = newFormat( "+++", toml.Unmarshal )
