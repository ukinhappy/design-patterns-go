package single

import (
	"math/rand"
	"sync"
	"time"
)

type Non struct {
	Value int
}

var non *Non
var once sync.Once

func GetNon() *Non {
	rand.Seed(time.Now().UnixNano())
	once.Do(func() {
		if non == nil {
			non = &Non{
				Value: rand.Intn(1000),
			}
		}
	})
	return non
}
