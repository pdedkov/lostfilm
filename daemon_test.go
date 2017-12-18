package lostfilm

import "testing"

func TestNewDaemon(t *testing.T) {
	_, err := NewDaemon("")
	if err == nil {
		t.Errorf("something went wrong %v", err)
	}
}