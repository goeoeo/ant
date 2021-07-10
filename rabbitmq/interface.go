package amqputil

import (
	"io"

	"github.com/streadway/amqp"
)

type AMQPConnPool interface {
	Acquire() (*amqp.Connection, error)
	Release(*amqp.Connection)
	io.Closer
}
