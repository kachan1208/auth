package cfg

import (
	"log"

	"github.com/kelseyhightower/envconfig"
)

type config struct {
	LogLevel          string `envconfig:"LOG_LEVEL" required:"true"`
	HTTPAddress       string `envconfig:"HTTP_ADDRESS" required:"true"`
	CassandraHost     string `envconfig:"CASSANDRA_HOST" required:"true"`
	CassandraPort     int    `envconfig:"CASSANDRA_PORT" required:"true"`
	CassandraLogin    string `envconfig:"CASSANDRA_LOGIN"`
	CassandraPassword string `envconfig:"CASSANDRA_PASSWORD"`
	CassandraKeyspace string `envconfig:"CASSANDRA_KEYSPACE" required:"true"`
}

var Config config

func init() {
	err := envconfig.Process("AUTH", &Config)
	if err != nil {
		log.Fatalf("Can't load config: %s ", err.Error())
	}
}
