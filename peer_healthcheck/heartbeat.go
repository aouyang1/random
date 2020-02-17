package main

import (
	"encoding/json"
	"net/http"
	"time"
)

type HeartBeat struct {
	Name string `json:"name"`
}

func heartbeat(w http.ResponseWriter, r *http.Request) {
	var hb HeartBeat
	if err := json.NewDecoder(r.Body).Decode(&hb); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	peers.Lock()
	if _, exists := peers.nodes[hb.Name]; !exists {
		logger.Printf("%s is not in the peer list\n", hb.Name)
		return
	}
	peers.nodes[hb.Name] = time.Now()
	peers.Unlock()
}
