package config

type Backend interface {
	Load(opts Options) error
	Unmarshal(path string, target any) error
	Raw() any
}
