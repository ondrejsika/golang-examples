run:
	go run .

generate-certs:
	openssl req -x509 -newkey rsa:2048 -keyout key.pem -out crt.pem -days 3650 -nodes -subj "/CN=example"
