package main

import (
	"context"
	"errors"
	"fmt"
)

type Subscriber struct {
	ctx         context.Context
	name        string
	subscribeCh chan string
}

func NewSubscriber(ctx context.Context) (*Subscriber, error) {
	if v := ctx.Value("name"); v != nil {
		subscriber := &Subscriber{ctx: ctx, name: v.(string), subscribeCh: make(chan string)}
		go subscriber.update()
		return subscriber, nil
	}
	return nil, errors.New("failed to create a new subscriber")
}

func (s *Subscriber) update() {
	for {
		select {
		case message := <-s.subscribeCh:
			fmt.Printf("%s is sent message as %s\n", s.name, message)
		case <-s.ctx.Done():
			wg.Done()
			return
		}
	}
}
