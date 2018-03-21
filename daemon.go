package lostfilm

import (
	"github.com/pdedkov/goconfig"
	"github.com/gen2brain/beeep"
	"os"
	"time"
)

// Config default daemon config
type config struct {
	Url     string        `toml:"url"`
	Timeout time.Duration `toml:"timeout"`
}

// Daemon struct with config
type daemon struct {
	config *config
}

// Run exec daemon
func (d *daemon) Run(quit <-chan os.Signal) error {
	ticker := time.NewTicker(d.config.Timeout * time.Minute)

	defer ticker.Stop()

	rssParser := newParser()

	last := time.Date(1970, time.December, 0, 0, 0, 0, 0, time.UTC)

	for {
		select {
		case <-ticker.C:
			items, err := rssParser.Parse(d.config.Url)
			if err != nil {
				return err
			}
			for _, item := range items {
				if item.PublishedParsed.After(last) {
					beeep.Notify("New episode", item.Title, "assets/information.png")
				}
			}
			last = time.Now()
		case <-quit:
			return nil
		}
	}
}

// NewDaemon createnew daemon
func NewDaemon(cfg string) (*daemon, error) {
	conf := &config{}
	err := goconfig.NewConfigFromFile(cfg, conf)
	if err != nil {
		return nil, err
	}
	return &daemon{conf}, nil
}
