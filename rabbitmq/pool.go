package amqputil

import (
	"errors"
	"io"
	"sync"

	"github.com/sirupsen/logrus"
	"github.com/streadway/amqp"
)

var (
	ErrConnectionClosed = errors.New("amqp: connection closed")
	ErrPoolClosed       = errors.New("rabbitmq: pool is closed")
)

type factoryFunc func() (io.Closer, error)

func NewConnPool(poolSize int, factory factoryFunc) AMQPConnPool {
	if poolSize <= 0 {
		poolSize = 1
	}
	return &pool{
		conns:   make(chan io.Closer, poolSize),
		factory: factory,
	}
}

type pool struct {
	m       sync.Mutex
	conns   chan io.Closer
	factory factoryFunc
	closed  bool
}

func (p *pool) Acquire() (*amqp.Connection, error) {
	select {
	case r, ok := <-p.conns:
		if !ok {
			return nil, ErrPoolClosed
		}
		conn, ok := r.(*amqp.Connection)
		if !ok || conn == nil || conn.IsClosed() {
			return nil, ErrConnectionClosed
		}

		return conn, nil
	default:
		r, err := p.factory()
		return r.(*amqp.Connection), err
	}
}

func (p *pool) Release(conn *amqp.Connection) {
	if conn == nil || conn.IsClosed() {
		return
	}

	p.m.Lock()
	defer p.m.Unlock()

	if p.closed {
		_ = conn.Close()
		return
	}

	select {
	case p.conns <- conn:
		logrus.Infof("Connection released")
	default:
		logrus.Infof("Pool is full, close connection ...")
		_ = conn.Close()
	}
}

func (p *pool) Close() error {
	p.m.Lock()
	defer p.m.Unlock()

	if p.closed {
		return nil
	}

	p.closed = true

	close(p.conns)
	for conn := range p.conns {
		_ = conn.Close()
		logrus.Infof("amqp connetion cloased ...")
	}

	logrus.Infof("Pool closed")

	return nil
}
