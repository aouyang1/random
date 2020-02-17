package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"strconv"
	"sync"
	"time"
)

type Peers struct {
	sync.Mutex
	nodes map[string]time.Time
}

func NewPeers() *Peers {
	return &Peers{nodes: make(map[string]time.Time)}
}

func (p *Peers) Reload(nodeList []string) {
	// convert node list into a map for faster lookup
	nodeMap := make(map[string]struct{})
	for _, n := range nodeList {
		if _, exists := nodeMap[n]; !exists {
			nodeMap[n] = struct{}{}
		}
	}

	p.Lock()
	defer p.Unlock()

	// evict nodes that aren't in the node list
	for n := range p.nodes {
		if _, exists := nodeMap[n]; !exists {
			delete(p.nodes, n)
		}
	}

	// add in nodes that are in the node list but not in the current peer list
	for n := range nodeMap {
		if _, exists := p.nodes[n]; !exists {
			p.nodes[n] = time.Time{}
		}
	}
}

func (p Peers) Alive(c *Config) []string {
	p.Lock()
	defer p.Unlock()
	lb := time.Now().Add(-c.AliveWindow)

	var alive []string
	for n, t := range p.nodes {
		if !t.Before(lb) {
			alive = append(alive, n)
		}
	}
	return alive
}

func (p Peers) Broadcast(c *Config) {
	ticker := time.NewTicker(c.BroadcastInterval)

	client := http.Client{Timeout: 10 * time.Second}
	body, err := json.Marshal(HeartBeat{Name: *node})
	if err != nil {
		logger.Fatalln(err)
	}
	port := strconv.Itoa(c.ListenPort)

	var wg sync.WaitGroup

	for {
		select {
		case <-ticker.C:
			// grab what is currently in the peers map
			p.Lock()
			nodeSnapshot := make([]string, 0, len(p.nodes))
			for k := range p.nodes {
				nodeSnapshot = append(nodeSnapshot, k)
			}
			p.Unlock()

			wg.Add(len(nodeSnapshot))
			for _, k := range nodeSnapshot {
				go func(n string) {
					defer wg.Done()
					req, err := http.NewRequest(
						"POST",
						"http://"+n+":"+port+"/heartbeat",
						bytes.NewBuffer(body),
					)
					if err != nil {
						logger.Println(err)
						return
					}
					_, err = client.Do(req)
					if err != nil {
						logger.Println(err)
					}
				}(k)
			}
			wg.Wait()
		}
	}
}
