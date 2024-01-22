FROM golang:1.21.6 as build
WORKDIR /build
COPY . .
RUN CGO_ENABLED=0 go build

FROM debian:12-slim
COPY \
  --from=build /build/ip_from_x_forwarded_for_server \
  /usr/local/bin/ip_from_x_forwarded_for_server
CMD ["/usr/local/bin/ip_from_x_forwarded_for_server"]
EXPOSE 8000
