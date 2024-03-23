package cachelru

import (
	"time"
)

var (
	capacity     int
	keyPairValue map[string]Item
)

const NOT_FOUND = -1

type Item struct {
	LastAcessDate time.Time
	Value         any
}

func NewCacheLRU(lenth int) {
	capacity = lenth
	keyPairValue = make(map[string]Item)
}

func GetCacheLRU(key string) any {

	value, ok := keyPairValue[key]

	if !ok {
		return NOT_FOUND
	}
	value.LastAcessDate = time.Now()

	return value.Value
}

func SetCacheLRU(key string, value any) {

	item := newItem(value)

	if len(keyPairValue) == capacity {
		delete(keyPairValue, getOldestItemKey())
	}

	keyPairValue[key] = item
}

func newItem(value any) Item {
	return Item{
		LastAcessDate: time.Now(),
		Value:         value,
	}
}

func getOldestItemKey() string {

	var (
		oldestKey string
	)

	for key, item := range keyPairValue {
		if oldestKey == "" {
			oldestKey = key
		}

		if item.LastAcessDate.Before(keyPairValue[oldestKey].LastAcessDate) {
			oldestKey = key
		}
	}

	return oldestKey
}
