package main

import (
	"context"
	"flag"
	"log"
	"os"
	"os/signal"
	"syscall"
	"wbl0/internal/apiserver"

	"github.com/BurntSushi/toml"
	"github.com/jackc/pgx/v4"
	"github.com/nats-io/stan.go"
)

func main() {
	var configPath string
	flag.StringVar(&configPath, "config-path", "configs/apiserver.toml", "path to config file")
	flag.Parse()

	config := apiserver.NewConfig()
	_, err := toml.DecodeFile(configPath, config)
	if err != nil {
		log.Fatal(err)
	}
	api := apiserver.Api{}

	api.DB, err = pgx.Connect(context.Background(), config.DatabaseURL)
	if err != nil {
		log.Println(err)
	}
	defer func(conn *pgx.Conn, ctx context.Context) {
		err := conn.Close(ctx)
		if err != nil {
			log.Printf("Error closing psql connection. %s", err)
		}
	}(api.DB, context.Background())

	err = api.CacheRecovery()
	if err != nil {
		log.Fatalln(err)
	}

	sc, err := stan.Connect(config.NatsClusterId, config.NatsClientId)
	if err != nil {
		log.Fatalln(err)
	}
	defer sc.Close()

	sc.Subscribe("foo", func(msg *stan.Msg) {
		api.ParseAndSaveMsg(msg)
	})

	server := api.StartHttp(config.Addr)

	log.Println("Start")
	log.Fatalln(server.ListenAndServe())

	signalChan := make(chan os.Signal, 1)
	cleanupDone := make(chan bool)
	signal.Notify(signalChan, os.Interrupt, syscall.SIGINT)
	go func() {
		for range signalChan {
			log.Printf("\nReceived an interrupt, closing connection...\n\n")
			sc.Close()
			cleanupDone <- true
		}
	}()
	<-cleanupDone
}
