package lru

import (
	"errors"
	"fmt"
)

type Cacher interface {
	Get(interface{}) (interface{}, error)
	Put(interface{}, interface{}) error
}

type lruCache struct {
	size      int
	remaining int
	cache     map[string]string
	queue     []string
}

func NewCache(size int) Cacher {
	return &lruCache{size: size, remaining: size, cache: make(map[string]string), queue: make([]string, 0)}
}

func (lru *lruCache) Get(key interface{}) (interface{}, error) {
	// Your code here....
	queue_length := len(lru.queue)
	key_str, key_ok := key.(string)
	val_str, val_ok := lru.cache[key_str]

	fmt.Println("key_str:", key_str, " val_str:", val_str)
	fmt.Println("key_ok:", key_ok, " val_ok:", val_ok)

	if key_ok && val_ok {
		if queue_length == 0 {
			return " ", errors.New("queue is empty")
		} else if queue_length > 0 && queue_length < 3 {
			lru.queue = append(lru.queue, key_str)
		} else if queue_length == 3 {
			lru.qDel(lru.queue[0])
			delete(lru.cache, lru.queue[0])
			lru.queue = append(lru.queue, key_str)
		}
	} else {
		return "", errors.New("key is not string type or not present in map")
	}
	return val_str, nil
}

func (lru *lruCache) Put(key, val interface{}) error {
	// Your code here....
	key_str, ok_key_str := key.(string)
	val_str, ok_val_str := val.(string)

	if len(lru.queue) < lru.size {
		if ok_key_str && ok_val_str {
			lru.cache[key_str] = val_str
			lru.queue = append(lru.queue, key_str)
		} else {
			return errors.New("key and val should be of string type")
		}
	} else if len(lru.queue) == lru.size {
		fmt.Println("deleting lru cache:", lru.queue[0])
		delete(lru.cache, lru.queue[0])
		lru.qDel(lru.queue[0])
		lru.queue = append(lru.queue, key_str)
		lru.cache[key_str] = val_str
		fmt.Println("queue after deleting:", lru.queue)
		fmt.Println("cache map after deleting:", lru.cache)
	}
	return nil
}

// Delete element from queue
func (lru *lruCache) qDel(ele string) {
	for i := 0; i < len(lru.queue); i++ {
		if lru.queue[i] == ele {
			oldlen := len(lru.queue)
			copy(lru.queue[i:], lru.queue[i+1:])
			lru.queue = lru.queue[:oldlen-1]
			break
		}
	}
}
