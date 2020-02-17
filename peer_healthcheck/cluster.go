package main

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

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
