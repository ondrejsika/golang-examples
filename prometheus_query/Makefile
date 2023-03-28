run:
	go run main.go

prometheus:
	prometheus

prometheus-delete-data:
	rm -rf ./data

loadgen:
	ab -n 100000 http://127.0.0.1:9090/metrics
