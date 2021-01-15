package usecase

import "sync"

type EventBus struct {
	subscribers []EventSubscriber
}

func NewEventBus(subscribers ...EventSubscriber) *EventBus {
	return &EventBus{subscribers: subscribers}
}

func (eb *EventBus) Append(s EventSubscriber) {
	eb.subscribers = append(eb.subscribers, s)
}

func (eb *EventBus) Receive(e Event) {
	wg := sync.WaitGroup{}
	for _, sub := range eb.subscribers {
		wg.Add(1)
		go func(sub EventSubscriber) {
			sub.Receive(e)
			wg.Done()
		}(sub)
	}
	wg.Wait()
}
