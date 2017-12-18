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
	var (
		config = flag.String("config",
			fmt.Sprintf("%s%s.lostfilm.toml", os.Getenv("HOME"), string(os.PathSeparator)),
			"config path",
		)
	)
	flag.Parse()
	d, err := lostfilm.NewDaemon(*config)
	if err != nil {
		log.Panic(err)
	}

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)

	d.Run(quit)
}
