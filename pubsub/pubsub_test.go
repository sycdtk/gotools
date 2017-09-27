package pubsub

import (
	"fmt"
	"strconv"
	"testing"
)

func TestSubscribe(t *testing.T) {
	c1 := &Client{Id: "client_1", Ip: "172.18.1.1"}
	c2 := &Client{Id: "client_2", Ip: "172.18.1.3"}
	c3 := &Client{Id: "client_3", Ip: "172.18.1.3"}

	srv := NewServer()

	srv.Subscribe(c1, "Topic")
	srv.Subscribe(c3, "Topic")

	srv.Subscribe(c1, "Topic2")
	srv.Subscribe(c2, "Topic2")

	for i := 0; i < 10; i++ {
		srv.PublishMessage("Topic", strconv.Itoa(i))
		srv.PublishMessage("Topic2", strconv.Itoa(i+10000000))
	}

	srv.Unsubscribe(c1, "Topic")

	fmt.Println("-----------------------------------------------------------")

	for i := 0; i < 10; i++ {
		srv.PublishMessage("Topic", strconv.Itoa(i))
		srv.PublishMessage("Topic2", strconv.Itoa(i+10000000))
	}
}
