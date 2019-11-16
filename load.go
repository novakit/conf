package conf

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

func LoadFile(dir string, name string, out interface{}) error {
	var err error
	for fmtName, f := range Formats {
		for _, ext := range f.Exts {
			filename := filepath.Join(dir, name+"."+ext)
			// check file existence
			if _, err = os.Stat(filename); err != nil {
				if os.IsNotExist(err) {
					continue
				} else {
					return fmt.Errorf("failed to check file '%s': %s", filename, err.Error())
				}
			}
			// load file content
			var buf []byte
			if buf, err = ioutil.ReadFile(filename); err != nil {
				return fmt.Errorf("failed to read file '%s': %s", filename, err.Error())
			}
			// convert to json
			if buf, err = f.ToJSON(buf); err != nil {
				return fmt.Errorf("failed to convert file '%s' from %s to JSON: %s", filename, fmtName, err.Error())
			}
			// unmarshal
			if err = json.Unmarshal(buf, out); err != nil {
				return fmt.Errorf("failed to unmarshal converted JSON from file '%s': %s", filename, err.Error())
			}
			// success
			return nil
		}
	}
	return fmt.Errorf("failed to find conf file '%s' with supported format in '%s'", name, dir)
}
