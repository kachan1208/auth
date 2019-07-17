package cfg

import (
	"log"

	"github.com/kelseyhightower/envconfig"
)

type config struct {
	LogLevel          string `envconfig:"LOG_LEVEL"`
	HTTPAddress       string `envconfig:"HTTP_ADDRESS`
	CassandraHost     string `envconfig:"CASSANDRA_HOST"`
	CassandraPort     string `envconfig:"CASSANDRA_PORT"`
	CassandraLogin    string `envconfig:"CASSANDRA_LOGIN"`
	CassandraPassword string `envconfig:"CASSANDRA_PASSWORD"`
}

var Config *config

func init() {
	err := envconfig.Process("AUTH", &Config)
	if err != nil {
		log.Fatalf("Can't load config %s ", err.Error())
	}
}
