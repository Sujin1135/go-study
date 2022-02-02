package main

import (
	"context"
	"fmt"
)

type Subscriber struct {
	ctx   context.Context
	name  string
	msgCh chan string
}

func NewSubscriber(name string, ctx context.Context) *Subscriber {
	return &Subscriber{ctx: ctx, name: name, msgCh: make(chan string)}
}

func (s *Subscriber) Subscribe(publisher *Publisher) {
	publisher.Subscribe(s.msgCh)
}

func (s *Subscriber) Update() {
	for {
		select {
		case msg := <-s.msgCh:
			fmt.Printf("publish message is %s to %s\n", msg, s.name)
		case <-s.ctx.Done():
			wg.Done()
			return
		}
	}
}
