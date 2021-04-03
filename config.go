package cbuff

import (
	"fmt"
	"github.com/caarlos0/env"
)

type CassandraConfig struct {
	Host     string `env:"CBUFF_CASSANDRA_HOST" envDefault:"localhost"`
	Port     int    `env:"CBUFF_CASSANDRA_PORT" envDefault:"9042"`
	Username string `env:"CBUFF_CASSANDRA_USER" envDefault:""`
	Password string `env:"CBUFF_CASSANDRA_PASSWORD" envDefault:""`
	Keyspace string `env:"CBUFF_CASSANDRA_KEYSPACE" envDefault:"frames"`

	MigrationsPath string
}

func NewCassandraConfigFromEnv() (*CassandraConfig, error) {
	config := new(CassandraConfig)

	err := env.Parse(config)
	if err != nil {
		return nil, err
	}

	return config, nil

}

func (config *CassandraConfig) ConnectionString() string {
	// cassandra://host:port/keyspace?username=value&param2=value2
	return fmt.Sprintf("cassandra://%s:%d/%s?username=%s&password=%s&x-multi-statement=true",
		config.Host, config.Port, config.Keyspace, config.Username, config.Password)
}
