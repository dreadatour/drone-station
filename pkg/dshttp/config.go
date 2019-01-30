package dshttp

import "fmt"

// Config is HTTP config
type Config struct {
	Host string `envconfig:"HOST" required:"true"`
	Port uint16 `envconfig:"PORT" required:"true"`
}

// Address returns HTTP address
func (c *Config) Address() string {
	return fmt.Sprintf("%s:%d", c.Host, c.Port)
}
