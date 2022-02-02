package main

import "context"

type Publisher struct {
	ctx         context.Context
	subscribeCh chan chan<- string // 일방향 채널
	publishCh   chan string
	subsribers  []chan<- string
}

func NewPublisher(ctx context.Context) *Publisher {
	return &Publisher{
		ctx:         ctx,
		subscribeCh: make(chan chan<- string),
		publishCh:   make(chan string),
		subsribers:  make([]chan<- string, 0),
	}
}

func (p *Publisher) Publish(msg string) {
	p.publishCh <- msg
}

func (p *Publisher) Update() {

}

func (p *Publisher) Subscribe() {

}

func main() {

}
