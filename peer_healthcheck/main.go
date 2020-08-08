package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
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

func main() {
	flag.Parse()

	if *node == "" {
		logger.Fatalln("No node name provided")
	}
	logger.Printf("Starting for node: %s\n", *node)

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
					logger.Println(err)
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
				fmt.Printf("Node: %s, %v\n", *node, peers.Alive(config))
			}
		}
	}()

	go peers.Broadcast(config)

	mux := http.NewServeMux()
	mux.Handle("/heartbeat", http.HandlerFunc(heartbeat))

	logger.Fatal(http.ListenAndServe(":"+strconv.Itoa(config.ListenPort), mux))
}
