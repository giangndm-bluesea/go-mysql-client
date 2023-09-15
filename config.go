package gomysqlclient

import "fmt"

type Config struct {
	Host     string
	Port     uint
	User     string
	Password string
}

func (c *Config) String() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%d)/", c.User, c.Password, c.Host, c.Port)
}
