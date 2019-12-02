package conf

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

type storageDir string

// Dir create a configuration storage from local directory
func Dir(dir string) Storage {
	return storageDir(dir)
}

func (d storageDir) Load(name string, out interface{}) error {
	var err error
	dir := string(d)
	for _, f := range Decoders {
		for _, ext := range f.IDs {
			filename := filepath.Join(dir, name+"."+ext)
			// check file existence
			if _, err = os.Stat(filename); err != nil {
				if os.IsNotExist(err) {
					continue
				} else {
					return fmt.Errorf("failed to check file existence '%s': %s", filename, err.Error())
				}
			}
			// load file content
			var buf []byte
			if buf, err = ioutil.ReadFile(filename); err != nil {
				return fmt.Errorf("failed to read file '%s': %s", filename, err.Error())
			}

			// decode
			if err = f.Decode(buf, out); err != nil {
				return fmt.Errorf("failed to decode file '%s': %s", filename, err.Error())
			}

			// success
			return nil
		}
	}
	return fmt.Errorf("failed to find conf '%s' from file with supported format in '%s'", name, dir)
}
