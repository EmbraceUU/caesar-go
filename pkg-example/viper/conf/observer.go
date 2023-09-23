package conf

type Observer interface {
	Update(config *Config)
}
