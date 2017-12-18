package lostfilm

import (
	"time"
	"github.com/pdedkov/goconfig"
	"os"
)

// Config default daemon config
type Config struct {
	Url string `toml:"url"`
	Timeout time.Duration `toml:"timeout"`
}

// Daemon struct with config
type Daemon struct {
	Config *Config
}

// Run exec daemon
func (d *Daemon) Run(quit <- chan os.Signal) error {
	ticker := time.NewTicker(d.Config.Timeout * time.Minute)

	defer ticker.Stop()

	parser := NewParser()
	for {
		select {
		case <- ticker.C:
			parser.Parse(d.Config.Url)
		case <- quit:
			return nil
		}
	}
}

// NewDaemon createnew daemon
func NewDaemon(config string) (*Daemon, error) {
	conf := &Config{}
	err := goconfig.NewConfigFromFile(config, conf)
	if err != nil {
		return nil, err
	}
	return &Daemon{conf}, nil
}

