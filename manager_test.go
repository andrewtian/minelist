package main

import (
	"log"
	"testing"
)

func TestListManager(t *testing.T) {
	svr := NewServer("asdf", "us.lichcraft.com:25565")
	if err := svr.Connect(); err != nil {
		t.Error("could not connect")
	}

	lm.AddServer(svr)
	lm.PingType(Active)

	if svr.LastPing() == nil {
		t.Error("last ping should not reutrn nil")
	}

	log.Println(svr.LastPing().latency)
}
