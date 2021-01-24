package random

import (
	"fmt"
	"sync"
)

type LRU struct {
	sync.Mutex

	size  int
	cache map[string]*value
	head  *value // most recently accessed or inserted value
	tail  *value // oldest that has been accessed or inserted
}

type value struct {
	k    string
	v    string
	prev *value // more recent
	next *value // older
}

func NewLRU(size int) *LRU {
	lru := &LRU{
		size:  size,
		cache: make(map[string]*value),
	}
	return lru
}

func (l *LRU) Put(k string, v string) {
	l.Lock()
	defer l.Unlock()
	val, exists := l.cache[k]
	if exists {
		// reuse existing value struct to save on gc. no change in length of cache.
		val.v = v
		if l.head != val {
			l.makeHead(val)
		}
		return
	}

	// we are exceeding the lru size so prune the tail
	if len(l.cache) == l.size {
		// update tail
		currTail := l.tail
		if currTail.prev != nil {
			l.tail = currTail.prev
			l.tail.next = nil
		}

		// remove the tail
		delete(l.cache, currTail.k)
	}

	val = &value{k: k, v: v}

	// insert
	l.cache[k] = val

	// update head
	// check if there is a head (first entry into lru)
	if l.head == nil {
		l.head = val
		l.tail = val
		return
	}

	l.makeHead(val)
	return

}

func (l *LRU) Get(k string) (string, error) {
	l.Lock()
	defer l.Unlock()
	val, exists := l.cache[k]
	if !exists {
		return "", fmt.Errorf("key %q does not exist", k)
	}

	if l.tail == val {
		if val.prev != nil {
			// update tail if the current value is the tail
			l.tail = val.prev
			l.tail.next = nil
		}
	}

	if l.head != val {
		// update elements around val
		val.prev.next = val.next
		if val.next != nil {
			// not at the tail
			val.next.prev = val.prev
		}

		l.makeHead(val)
	}

	return val.v, nil
}

func (l *LRU) Size() int {
	l.Lock()
	defer l.Unlock()

	return len(l.cache)
}

func (l *LRU) Elements() []string {
	l.Lock()
	defer l.Unlock()

	v := l.head
	elem := make([]string, 0, len(l.cache))
	for {
		if v == nil {
			break
		}
		elem = append(elem, v.v)
		fmt.Println(elem)
		v = v.next
	}
	return elem
}

func (l *LRU) makeHead(v *value) {
	l.head.prev = v
	v.next = l.head
	v.prev = nil
	l.head = v
}
