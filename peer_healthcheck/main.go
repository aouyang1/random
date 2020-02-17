package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
	"sync"
	"time"
)

var (
	logger = log.New(os.Stderr, "peer-healthcheck: ", log.Ldate|log.Ltime|log.Lshortfile)

	peers  *Peers
	config *Config

	node = flag.String("node", "", "node name")
)

type Config struct {
	AliveWindow           time.Duration
	BroadcastInterval     time.Duration
	ConfigRefreshInterval time.Duration
	ListenPort            int
}

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

func (p Peers) Alive(dur time.Duration) []string {
	p.Lock()
	defer p.Unlock()
	lb := time.Now().Add(-dur)

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
		log.Fatalln(err)
	}

	var wg sync.WaitGroup

	for {
		select {
		case <-ticker.C:
			// grab what is currently in the peers map
			peers.Lock()
			nodeSnapshot := make([]string, 0, len(peers.nodes))
			for k := range p.nodes {
				nodeSnapshot = append(nodeSnapshot, k)
			}
			peers.Unlock()

			wg.Add(len(nodeSnapshot))
			for _, k := range nodeSnapshot {
				go func(n string) {
					defer wg.Done()
					req, err := http.NewRequest(
						"POST",
						"http://"+n+":"+strconv.Itoa(c.ListenPort)+"/heartbeat",
						bytes.NewBuffer(body),
					)
					if err != nil {
						log.Println(err)
						return
					}
					_, err = client.Do(req)
					if err != nil {
						log.Println(err)
					}
				}(k)
			}
			wg.Wait()
		}
	}
}

type HeartBeat struct {
	Name string `json:"name"`
}

func heartbeat(w http.ResponseWriter, r *http.Request) {
	var hb HeartBeat
	if err := json.NewDecoder(r.Body).Decode(&hb); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if _, exists := peers.nodes[hb.Name]; !exists {
		logger.Printf("%s is not in the peer list\n", hb.Name)
		return
	}

	ct := time.Now()

	peers.Lock()
	peers.nodes[hb.Name] = ct
	peers.Unlock()
}

type Cluster struct {
	Cluster []string `json:"cluster"`
}

func loadCluster() ([]string, error) {
	f, err := os.Open("cluster.json")
	if err != nil {
		return nil, err
	}
	defer f.Close()
	bytes, err := ioutil.ReadAll(f)
	if err != nil {
		return nil, err
	}
	var c Cluster
	if err := json.Unmarshal(bytes, &c); err != nil {
		return nil, err
	}

	return c.Cluster, nil
}

func main() {
	flag.Parse()

	if *node == "" {
		log.Fatalln("No node name provided")
	}
	log.Printf("Starting for node: %s\n", *node)

	config = &Config{
		AliveWindow:           30 * time.Second,
		BroadcastInterval:     5 * time.Second,
		ConfigRefreshInterval: 10 * time.Second,
		ListenPort:            1212,
	}

	peers = NewPeers()

	nodes, err := loadCluster()
	if err != nil {
		panic(err)
	}

	peers.Reload(nodes)

	go func() {
		ticker := time.NewTicker(config.ConfigRefreshInterval)
		for {
			select {
			case <-ticker.C:
				nodes, err := loadCluster()
				if err != nil {
					log.Println(err)
					continue
				}

				peers.Reload(nodes)
			}
		}
	}()

	// show what's currently marked as alive
	go func() {
		ticker := time.NewTicker(3 * time.Second)
		for {
			select {
			case <-ticker.C:
				fmt.Printf("Node: %s, %v\n", *node, peers.Alive(config.AliveWindow))
			}
		}
	}()

	go peers.Broadcast(config)

	mux := http.NewServeMux()
	mux.Handle("/heartbeat", http.HandlerFunc(heartbeat))

	log.Fatal(http.ListenAndServe(":"+strconv.Itoa(config.ListenPort), mux))
}
