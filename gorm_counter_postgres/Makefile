run:
	docker rm -f postgres || true
	docker run -d --name postgres -p 5432:5432 -e POSTGRES_PASSWORD=pg postgres:16
	slu wait-for tcp -a 127.0.0.1:5432
	go run main.go
	go run main.go
	go run main.go
	docker rm -f postgres
