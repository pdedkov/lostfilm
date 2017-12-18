package lostfilm

import "github.com/mmcdole/gofeed"

type Parser struct {
	Parser *gofeed.Parser
}

// Parse parse rss by it url
func (p *Parser) Parse(url string) (string, error) {
	feed, err := p.Parser.ParseURL(url)
	if err != nil {
		return "", err
	}
	return feed.Title, nil
}

// NewParser return new parser instance
func NewParser() *Parser {
	return &Parser{gofeed.NewParser()}
}

