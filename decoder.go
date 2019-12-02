package conf

import (
	"encoding/json"
	"fmt"
	"github.com/BurntSushi/toml"
	"github.com/creasty/defaults"
	"gopkg.in/yaml.v2"
)

type Decoder struct {
	IDs    []string
	ToJSON func(content []byte) (out []byte, err error)
}

func (f Decoder) Decode(buf []byte, out interface{}) error {
	var err error
	// convert to json
	if buf, err = f.ToJSON(buf); err != nil {
		return fmt.Errorf("failed to convert: %s", err.Error())
	}
	// unmarshal
	if err = json.Unmarshal(buf, out); err != nil {
		return fmt.Errorf("failed to unmarshal: %s", err.Error())
	}
	// set defaults
	if err = defaults.Set(out); err != nil {
		return fmt.Errorf("failed to apply defaults: %s", err.Error())
	}
	return nil
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
