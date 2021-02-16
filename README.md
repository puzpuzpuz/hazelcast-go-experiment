# Turbine Go Experiment

Simple experiment aimed towards understanding memory footprint of HTTP server in combination with Hazelcast Go client.

Requires Go v1.13.

## Start

First of all, you need to start Hazelcast IMDG cluster. The simplest way is to start a single member cluster with Docker:
```bash
docker run --net=host hazelcast/hazelcast:3.12.12
```

With Go:
```bash
go get -d github.com/gorilla/mux
go get -d github.com/hazelcast/hazelcast-go-client
go run main.go
```

## Test

```bash
curl -X POST -H "Content-Type: text/plain" -d '{"foo": "bar"}' http://127.0.0.1:8080/api/testmap/testkey
```

## Load Test

The following command with be generating 1,000 RPS load for one minute:
```bash
npx autocannon -c 10 -d 60 -R 1000 -m POST -b "testvalue" -H "Content-Type: text/plain" http://localhost:8080/api/testmap/testkey
```
