package main

import (
	"awesomeProject/internal/base-module/bootstrap"
	"context"
	"flag"
	"github.com/labstack/gommon/log"
)

// Клиент, запрашивающий у сервера данные
func main() {
	var (
		configPath string
		etcdAddr   string
	)

	flag.StringVar(&configPath, "config", "", "path to config")
	flag.StringVar(&etcdAddr, "etcd", "", "etcd3 address")
	flag.Parse()

	ctx := context.Background()

	err := bootstrap.Run(ctx, &bootstrap.Config{
		ETCDAddr:   etcdAddr,
		ConfigPath: configPath,
	})

	//log.Sync() из какого пакета и зачем оно надо
	if err != nil {
		log.Fatal(err)
	}
	log.Info("awesome project stopped")
}
