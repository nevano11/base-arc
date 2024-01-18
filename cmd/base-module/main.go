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
		configPath     string
		configFilename string
		etcdAddr       string
	)

	flag.StringVar(&configPath, "config", "", "path to config")
	flag.StringVar(&configFilename, "configFile", "", "config filename")
	flag.StringVar(&etcdAddr, "etcd", "", "etcd address")
	flag.Parse()

	if len(configPath) > 0 {
		log.Infof("Client started with config=%s\n", configPath)
	}
	if len(etcdAddr) > 0 {
		log.Infof("Client started with etcd=%s\n", etcdAddr)
	}

	ctx := context.Background()

	err := bootstrap.Run(ctx, &bootstrap.PathConfig{
		ETCDAddr:       etcdAddr,
		ConfigPath:     configPath,
		ConfigFilename: configFilename,
	})

	//log.Sync()
	if err != nil {
		log.Fatal(err)
	}
	log.Info("awesome project client stopped")
}
