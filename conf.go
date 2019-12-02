package conf

// DefaultStorage default conf storage, from "conf" directory
var DefaultStorage = Dir("conf")

// Load shortcut to DefaultStorage.Load
func Load(name string, out interface{}) error {
	return DefaultStorage.Load(name, out)
}
