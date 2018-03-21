package lostfilm

import "github.com/mmcdole/gofeed"

type parser struct {
	parser *gofeed.Parser
}

// Parse parse rss by it url
func (p *parser) Parse(url string) ([]*gofeed.Item, error) {
	feed, err := p.parser.ParseURL(url)
	if err != nil {
		return nil, err
	}

	return feed.Items, nil
}

// NewParser return new parser instance
func newParser() *parser {
	return &parser{gofeed.NewParser()}
}
