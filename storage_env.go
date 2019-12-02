package conf

import (
	"encoding/json"
	"fmt"
	"github.com/creasty/defaults"
	"os"
	"strings"
)

type storageEnv string

// Env create a configuration storage from env prefix
func Env(prefix string) Storage {
	return storageEnv(prefix)
}

func (d storageEnv) Load(name string, out interface{}) error {
	pfx := string(d)
	var err error
	for fmtName, f := range Decoders {
		for _, ext := range f.IDs {
			key := strings.ToUpper(pfx) + "_" + strings.ToUpper(name) + "_" + strings.ToUpper(ext)
			val := os.Getenv(key)
			if len(val) == 0 {
				continue
			}
			buf := []byte(val)
			// convert to json
			if buf, err = f.ToJSON(buf); err != nil {
				return fmt.Errorf("failed to convert env '%s' from %s to JSON: %s", key, fmtName, err.Error())
			}
			// unmarshal
			if err = json.Unmarshal(buf, out); err != nil {
				return fmt.Errorf("failed to unmarshal converted JSON from env '%s': %s", key, err.Error())
			}
			// set defaults
			if err = defaults.Set(out); err != nil {
				return fmt.Errorf("failed to set defaults to env '%s': %s", key, err.Error())
			}
			// success
			return nil
		}
	}
	return fmt.Errorf("failed to find conf env '%s' with supported format", name)
}
