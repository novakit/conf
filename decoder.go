package conf

import (
	"encoding/json"
	"github.com/BurntSushi/toml"
	"gopkg.in/yaml.v2"
)

type Decoder struct {
	IDs    []string
	ToJSON func(content []byte) (out []byte, err error)
}

var Decoders = map[string]Decoder{
	"JSON": {
		IDs: []string{"json"},
		ToJSON: func(content []byte) ([]byte, error) {
			return content, nil
		},
	},
	"YAML": {
		IDs: []string{"yml", "yaml"},
		ToJSON: func(content []byte) (bytes []byte, err error) {
			var m map[string]interface{}
			if err = yaml.Unmarshal(content, &m); err != nil {
				return
			}
			return json.Marshal(m)
		},
	},
	"TOML": {
		IDs: []string{"toml"},
		ToJSON: func(content []byte) (out []byte, err error) {
			var m map[string]interface{}
			if err = toml.Unmarshal(content, &m); err != nil {
				return
			}
			return json.Marshal(m)
		},
	},
}
