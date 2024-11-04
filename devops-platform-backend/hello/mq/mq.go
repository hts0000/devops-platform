package mq

import "context"

type Publisher interface {
	Publish(context.Context, any) error
}

type Subscriber interface {
	Subscribe(context.Context) (ch chan any, cleanUp func(), err error)
}
