package pool

import (
	"sync/atomic"
	"testing"
)

var idx int32

type TestConnection struct {
	id int32
}

func (conn *TestConnection) Ping() error {
	return nil
}
func (conn *TestConnection) Close() error {
	return nil
}

func product(args ...interface{}) (Connection, error) {
	c := new(TestConnection)
	c.id = atomic.AddInt32(&idx, 1)
	return c, nil
}

func TestDefaultPool(t *testing.T) {
	dp, err := NewDefaultPool(DefaultConfig(), product)
	if err != nil {
		t.Fatal("create default pool fail", err)
	}

	for i := 0; i < 10; i++ {
		t.Run("acquire and release", func(t *testing.T) {
			c, err := dp.Acquire()
			if err != nil {
				t.Error("acquire fail", err)
			} else {
				t.Logf("connection id[%d] acquired", c.(*TestConnection).id)
				if i%2 == 0 {
					dp.Release(c)
					t.Logf("connection id[%d] release", c.(*TestConnection).id)
				}
			}
		})
	}

	dp.Shutdown()
}
