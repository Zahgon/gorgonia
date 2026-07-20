package xvm

import (
	"context"
	"sync"

	"gorgonia.org/gorgonia"
)

type publisher struct {
	id          int64
	publisher   <-chan gorgonia.Value
	subscribers []chan<- gorgonia.Value
}

type subscriber struct {
	id         int64
	publishers []<-chan gorgonia.Value
	subscriber chan<- ioValue
}

type pubsub struct {
	publishers  []*publisher
	subscribers []*subscriber
}

func (p *pubsub) run(ctx context.Context) (context.CancelFunc, *sync.WaitGroup) {
	_ = "STUB: not implemented"
	return *new(context.CancelFunc), nil
}

func merge(ctx context.Context, globalWG *sync.WaitGroup, out chan<- ioValue, cs ...<-chan gorgonia.Value) {
	_ = "STUB: not implemented"
	return
}

func broadcast(ctx context.Context, globalWG *sync.WaitGroup, ch <-chan gorgonia.Value, cs ...chan<- gorgonia.Value) {
	_ = "STUB: not implemented"
	return
}
