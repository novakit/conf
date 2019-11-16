package conf

import "fmt"

type Loader struct {
	Name   string
	Target interface{}
	Loaded func()
}

var loaders []*Loader

func RegisterLoader(l *Loader) {
	loaders = append(loaders, l)
}

func RunLoaders(dir string) error {
	for _, l := range loaders {
		if err := LoadFile(dir, l.Name, l.Target); err != nil {
			return fmt.Errorf("failed to run conf loader for ‘%s’: %s", l.Name, err.Error())
		}
		if l.Loaded != nil {
			l.Loaded()
		}
	}
	return nil
}
