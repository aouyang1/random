FROM golang:1.13.5
RUN mkdir /go/src/app
WORKDIR /go/src/app
CMD go run main.go heartbeat.go cluster.go peers.go -node $NODE
