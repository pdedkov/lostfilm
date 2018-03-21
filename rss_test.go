package lostfilm

import "testing"

func TestNewRssParser(t *testing.T) {
	rss := newParser()
	if rss.parser == nil {
		t.Error("wrong parser")
	}
}
