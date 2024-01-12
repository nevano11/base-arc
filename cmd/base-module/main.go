package main

import (
	"awesomeProject/internal/base-module/bootstrap"
	"context"
	"flag"
	"github.com/labstack/gommon/log"
)

// Client requesting data from the server
func main() {
	var (
		configPath string
		etcdAddr   string
	)

	flag.StringVar(&configPath, "config", "", "path to config")
	flag.StringVar(&etcdAddr, "etcd", "", "etcd3 address")
	flag.Parse()

	if len(configPath) > 0 {
		log.Infof("Client started with config=%s\n", configPath)
	}
	if len(etcdAddr) > 0 {
		log.Infof("Client started with etcd=%s\n", etcdAddr)
	}

	ctx := context.Background()

	err := bootstrap.Run(ctx, &bootstrap.Config{
		ETCDAddr:   etcdAddr,
		ConfigPath: configPath,
	})

	//log.Sync()
	if err != nil {
		log.Fatal(err)
	}
	log.Info("awesome project client stopped")
}
