FROM golang:1.19 as buil
WORKDIR /build
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 go build

FROM debian:11-slim
COPY --from=buil /build/counter_api_server_redis /
CMD ["/counter_api_server_redis"]
EXPOSE 8000
