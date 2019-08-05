package main

import (
	"net/http"
	"os"

	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	"github.com/gocql/gocql"

	"github.com/kachan1208/auth/src/cfg"
	"github.com/kachan1208/auth/src/controller"
	"github.com/kachan1208/auth/src/dao"
	transport "github.com/kachan1208/auth/src/transport/http"
)

func main() {
	logger := initLogger()
	level.Info(logger).Log("msg", "Service starting")

	sess := initSess()
	defer sess.Close()

	level.Info(logger).Log("msg", "DB session created")

	repo := dao.NewTokenRepo(sess)
	controller := controller.NewController(repo)
	handler := transport.NewHandler(cfg.Config.HTTPAddress, logger, controller)

	server := &http.Server{
		Addr:    handler.Address,
		Handler: handler.Router,
	}

	errs := make(chan error)
	go func() {
		if err := server.ListenAndServe(); err != nil {
			level.Error(logger).Log(err)
			errs <- err
		}
	}()

	level.Info(logger).Log("msg", "Service started", "address:", cfg.Config.HTTPAddress)
	level.Error(logger).Log("exit", <-errs)
}

func initSess() *gocql.Session {
	cluster := gocql.NewCluster(cfg.Config.CassandraHost)
	if cfg.Config.CassandraLogin != "" && cfg.Config.CassandraPassword != "" {
		cluster.Authenticator = gocql.PasswordAuthenticator{cfg.Config.CassandraLogin, cfg.Config.CassandraPassword}
	}

	cluster.Keyspace = cfg.Config.CassandraKeyspace
	cluster.Consistency = gocql.LocalQuorum
	cluster.Port = cfg.Config.CassandraPort
	sess, err := cluster.CreateSession()
	if err != nil {
		panic(err)
	}

	return sess
}
func initLogger() log.Logger {
	var logger log.Logger

	logger = log.NewLogfmtLogger(os.Stdout)
	logger = log.NewSyncLogger(logger)
	logger = level.NewFilter(logger, level.AllowDebug())
	logger = log.With(logger,
		"service", "auth",
		"ts", log.DefaultTimestampUTC,
		"caller", log.DefaultCaller,
	)

	return logger
}
