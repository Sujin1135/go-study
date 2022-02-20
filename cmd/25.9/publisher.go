package main

import "context"

type Publisher struct {
	ctx         context.Context
	publishCh   chan string
	subscribeCh chan chan<- string
	subscribers []chan<- string
}

func NewPublisher(ctx context.Context) *Publisher {
	publisher := &Publisher{
		ctx:         ctx,
		subscribeCh: make(chan chan<- string),
		subscribers: make([]chan<- string, 0),
		publishCh:   make(chan string),
	}
	go publisher.update()
	return publisher
}

func (p *Publisher) Subscribe(ch chan<- string) {
	p.subscribeCh <- ch
}

func (p *Publisher) Publish(message string) {
	p.publishCh <- message
}

func (p *Publisher) update() {
	for {
		select {
		case ch := <-p.subscribeCh:
			p.subscribers = append(p.subscribers, ch)
		case message := <-p.publishCh:
			for _, ch := range p.subscribers {
				ch <- message
			}
		case <-p.ctx.Done():
			wg.Done()
			return
		}
	}
}
