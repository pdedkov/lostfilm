package lostfilm

import "testing"

func TestNewRssParser(t *testing.T) {
	rss := NewParser()
	if rss.Parser == nil {
		t.Error("wrong parser")
	}
}
