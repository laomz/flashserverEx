package common

import "sync"

var G_ObjectPoolTest = sync.Pool{
	New: func() interface{} {
		return struct{}{}
	},
}
