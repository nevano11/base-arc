package main

import (
	"awesomeProject/internal/core-module/bootstrap"
	"context"
	"flag"
	"github.com/labstack/gommon/log"
)

// A server that accesses Tarantool and gives data from it to the client
func main() {
	var (
		configPath     string
		configFilename string
		etcdAddr       string
	)

	flag.StringVar(&configPath, "config", "", "path to config")
	flag.StringVar(&configFilename, "configFile", "", "config filename")
	flag.StringVar(&etcdAddr, "etcd", "", "etcd address")
	flag.Parse()

	if len(configPath) > 0 {
		log.Infof("Server started with config=%s\n", configPath)
	}
	if len(configFilename) > 0 {
		log.Infof("Server started with configFilename=%s\n", configFilename)
	}
	if len(etcdAddr) > 0 {
		log.Infof("Server started with etcd=%s\n", etcdAddr)
	}

	ctx := context.Background()

	err := bootstrap.Run(ctx, &bootstrap.PathConfig{
		ETCDAddr:       etcdAddr,
		ConfigPath:     configPath,
		ConfigFilename: configFilename,
	})

	if err != nil {
		log.Fatal(err)
	}
	log.Info("awesome project server stopped")
}
