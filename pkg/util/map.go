package util

import "sync"

type SyncMap[k comparable, v any] struct {
	m sync.Map
}

// Store 存储键值对
func (tsm *SyncMap[K, V]) Store(key K, value V) {
	tsm.m.Store(key, value)
}

// Load 获取键对应的值
func (tsm *SyncMap[K, V]) Load(key K) (V, bool) {
	val, ok := tsm.m.Load(key)
	if !ok {
		var zero V
		return zero, false
	}
	return val.(V), true
}

// Delete 删除键值对
func (tsm *SyncMap[K, V]) Delete(key K) {
	tsm.m.Delete(key)
}

func (tsm *SyncMap[K, V]) Range(f func(K, V) bool) {
	tsm.m.Range(func(key, value any) bool {
		return f(key.(K), value.(V))
	})
}
