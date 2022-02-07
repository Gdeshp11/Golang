package lru

import (
	"errors"
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

	//check if key is string and value is found in map
	if key_ok && val_ok {
		if queue_length == 0 {
			return " ", errors.New("queue is empty") //throw error if queue is empty
		} else if queue_length > 0 && queue_length < lru.size {
			lru.queue = append(lru.queue, key_str) //append the key used in 'Get' to queue
		} else if queue_length == lru.size {
			lru.qDel(lru.queue[0]) //delete first element in queue when queue is already full
			delete(lru.cache, lru.queue[0])
			lru.queue = append(lru.queue, key_str) // append the key used in 'Get' to queue
		}
	} else {
		return "", errors.New("key is not string type or not present in map")
	}
	return val_str, nil //return value of requested key and nil for success
}

func (lru *lruCache) Put(key, val interface{}) error {
	// Your code here....
	key_str, ok_key_str := key.(string)
	val_str, ok_val_str := val.(string)

	if len(lru.queue) < lru.size {
		// check if key and val is of string type
		if ok_key_str && ok_val_str {
			lru.cache[key_str] = val_str           //insert key val pair in map
			lru.queue = append(lru.queue, key_str) //append key in queue
		} else {
			return errors.New("key and val should be of string type")
		}
	} else if len(lru.queue) == lru.size { //when queue if full
		delete(lru.cache, lru.queue[0])        // delete key val pair in map
		lru.qDel(lru.queue[0])                 //delete first key in queue if its full
		lru.queue = append(lru.queue, key_str) //append the new key in queue
		lru.cache[key_str] = val_str           //update map with new key val
	}
	return nil //return nil if successful
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
