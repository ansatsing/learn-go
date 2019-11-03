package ch15_objectPool

import (
	"errors"
	"time"
)

type connectionPool struct {
	chn chan *connection
	len int
	cap int
}

func NewConnectionPool(len int, cap int) *connectionPool {
	conPool := new(connectionPool)
	conPool.chn = make(chan *connection, cap)
	conPool.len = len
	conPool.cap = cap
	for i := 0; i < len; i++ {
		conPool.chn <- new(connection)
	}
	return conPool
}

func (conPool *connectionPool) getConnection(timOut time.Duration) (*connection, error) {
	if conPool.len == 0 {
		return nil, errors.New("the pool is empty")
	}
	select {
	case conn := <-conPool.chn:
		conPool.len--
		return conn, nil
	case <-time.After(timOut):
		return nil, errors.New("get connection from pool time out")
	}
}

func (conPool *connectionPool) releaseConnection(conn *connection) error {
	if conPool.len == conPool.cap {
		return errors.New("the pool is overflow")
	}
	select {
	case conPool.chn <- conn:
		return nil
	default:
		return errors.New("overflow")
	}
}
