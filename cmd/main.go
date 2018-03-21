package main

import (
	"flag"
	"fmt"
	"github.com/pdedkov/lostfilm"
	"os"
	"log"
	"os/signal"
)

func main() {
	currentDir, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	var (
		config = flag.String("config", fmt.Sprintf("%s%sconfig.toml", currentDir, string(os.PathSeparator)), "config path")
	)

	flag.Parse()
	d, err := lostfilm.NewDaemon(*config)
	if err != nil {
		log.Panic(err)
	}

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)

	log.Fatal(d.Run(quit))
}
