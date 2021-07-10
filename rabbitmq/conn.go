package amqputil

import (
	"sync"
	"time"

	"github.com/sirupsen/logrus"
	"github.com/streadway/amqp"
)

type AmqpConn struct {
	conn             *amqp.Connection
	reconnectSuccess chan struct{}
	amqpConnPool     AMQPConnPool
	once             sync.Once

	intervalSeconds int //重试间隔 单位:秒

	quit chan struct{}
}

func NewAmqpConn(amqpConnPool AMQPConnPool, quit chan struct{}) *AmqpConn {
	return &AmqpConn{
		amqpConnPool:    amqpConnPool,
		once:            sync.Once{},
		quit:            quit,
		intervalSeconds: 5,
	}
}

func (c *AmqpConn) Conn() *amqp.Connection {
	c.once.Do(func() {
		c.conn = c.mustConnect()

		go c.reconnect()
		go c.release()
	})

	if !c.conn.IsClosed() {
		return c.conn
	}

	//连接不可以用，等待重连
	<-c.reconnectSuccess

	return c.conn

}

func (c *AmqpConn) Close() {
	//通知重连程序退出
	c.quit <- struct{}{}
	if c.conn != nil && !c.conn.IsClosed() {
		c.amqpConnPool.Release(c.conn)
	}
	c.conn = nil

	return
}

//release
func (c *AmqpConn) release() {
	<-c.quit

	if c.conn == nil {
		return
	}

	if !c.conn.IsClosed() {
		c.amqpConnPool.Release(c.conn)
	}
	c.conn = nil
}

//重连
func (c *AmqpConn) reconnect() {
	c.reconnectSuccess = make(chan struct{})
	closeNotify := make(chan *amqp.Error)
	c.conn.NotifyClose(closeNotify)

	for {
		select {
		case err, ok := <-closeNotify:
			if !ok {
				return
			}

			if err != nil {
				logrus.Error("AmqpConn.reconnect NotifyClose: ", err)

				if c.conn != nil && !c.conn.IsClosed() {
					_ = c.conn.Close()
					close(closeNotify)
				}

				//新建连接
				c.conn = c.mustConnect()

				//获取连接成功
				close(c.reconnectSuccess)
				go c.reconnect()
			}

		case <-c.quit:
			return
		}
	}
}

func (c *AmqpConn) mustConnect() (conn *amqp.Connection) {
	var (
		err error
	)

	for {
		if conn, err = c.amqpConnPool.Acquire(); err == nil {
			return
		}

		logrus.Infof("AmqpConn.mustConnect retry acquire after  %d seconds, error:%s", c.intervalSeconds, err)
		time.Sleep(time.Duration(c.intervalSeconds) * time.Second)
	}
}
