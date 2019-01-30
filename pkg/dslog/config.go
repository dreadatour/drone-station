package dslog

// Config is logging config
type Config struct {
	Level  string `envconfig:"LEVEL" default:"debug"`
	Format string `envconfig:"FORMAT" default:"text"`
}
