package patterns

import "sync"

var once sync.Once

type singleton struct {
	Text string
}

var (
	instance singleton
)

func GetSingletonInstance(text string) singleton {

	once.Do(func() {
		instance = singleton{text}
	})

	return instance
}