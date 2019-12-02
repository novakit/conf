package conf

// Storage configuration storage, could be a local directory, a etcd endpoint or something else
type Storage interface {
	// Load load specified configuration into interface, tag 'json' and 'default' are supported
	Load(name string, out interface{}) error
}
