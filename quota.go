package random

import (
	"fmt"
	"sync"
	"time"
)

type manager struct {
	mu                 sync.Mutex
	rules              map[string]*rule
	updateRate         time.Duration
	lastUpdateDuration time.Duration
}

func newManager() *manager {
	return &manager{
		rules:      make(map[string]*rule),
		updateRate: 1 * time.Second,
	}
}

func (m *manager) addRule(key string, r *rule) {
	m.mu.Lock()
	m.rules[key] = r
	m.mu.Unlock()
}

func (m *manager) run() {
	ticker := time.NewTicker(m.updateRate)
	go func() {
		for {
			select {
			case <-ticker.C:
				m.addTokens()
			}
		}
	}()
}

func (m *manager) addTokens() {
	m.mu.Lock()
	start := time.Now()
	for _, r := range m.rules {
		r.addToken(m.updateRate)
	}
	m.lastUpdateDuration = time.Now().Sub(start)
	m.mu.Unlock()
}

func (m *manager) useToken(key string) (bool, error) {
	m.mu.Lock()
	var ok bool
	var err error
	if r, exists := m.rules[key]; exists {
		ok = r.useToken()
	} else {
		err = fmt.Errorf("key %s does not exist as a quota rule", key)
	}
	m.mu.Unlock()
	return ok, err
}

type rule struct {
	qps        int
	window     time.Duration
	count      int
	maxQueries int
}

func newRule(qps int, window time.Duration) *rule {
	return &rule{
		qps:        qps,
		window:     window,
		count:      0,
		maxQueries: int(window.Seconds() * float64(qps)),
	}
}

func (r *rule) addToken(dur time.Duration) {
	numTokens := int(dur.Seconds() * float64(r.qps))

	if r.count+numTokens > r.maxQueries {
		r.count = r.maxQueries
	} else {
		r.count += numTokens
	}
}

func (r *rule) useToken() bool {
	if r.count == 0 {
		return false
	}
	r.count--
	return true
}
