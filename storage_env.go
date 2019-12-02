package conf

import (
	"fmt"
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
	for _, f := range Decoders {
		for _, ext := range f.IDs {
			// retrieve value
			key := strings.ToUpper(pfx) + "_" + strings.ToUpper(name) + "_" + strings.ToUpper(ext)
			val := []byte(os.Getenv(key))
			if len(val) == 0 {
				continue
			}

			// decode
			if err = f.Decode(val, out); err != nil {
				return fmt.Errorf("failed to decode key '%s': %s", key, err.Error())
			}

			// success
			return nil
		}
	}
	return fmt.Errorf("failed to find conf '%s' from env with supported format and prefix '%s'", name, pfx)
}
