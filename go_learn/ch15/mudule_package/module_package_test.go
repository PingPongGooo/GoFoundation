package mudule_package

import (
	cm "github.com/easierway/concurrent_map"
	"testing"
)

func TestConcurrenMap(t *testing.T)  {
	m := cm.CreateConcurrentMap(99)
	m.Set(cm.StrKey("key"),33)
	t.Log(m.Get(cm.StrKey("key")))
}
