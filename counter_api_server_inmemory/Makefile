dev:
	go run main.go

IMAGE = ondrejsika/counter-api-server-inmemory

docker-build-and-push:
	make docker-build
	make docker-push

docker-build:
	docker build -t ${IMAGE} .

docker-push:
	docker push ${IMAGE}

docker-run:
	docker run -p 8000:8000 ${IMAGE}
