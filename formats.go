package conf

import (
	"encoding/json"
	"gopkg.in/yaml.v2"
)

type Format struct {
	Exts   []string
	ToJSON func(content []byte) ([]byte, error)
}

var (
	Formats = map[string]Format{
		"JSON": {
			Exts: []string{"json"},
			ToJSON: func(content []byte) ([]byte, error) {
				return content, nil
			},
		},
		"YAML": {
			Exts: []string{"yml", "yaml"},
			ToJSON: func(content []byte) (bytes []byte, err error) {
				var m map[string]interface{}
				if err = yaml.Unmarshal(content, &m); err != nil {
					return
				}
				return json.Marshal(m)
			},
		},
	}
)
