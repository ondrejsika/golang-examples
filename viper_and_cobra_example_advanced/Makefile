all:
	@make test-no-config
	@make test-file
	@make test-env
	@make test-flag

test-no-config:
	go run main.go hello || true

test-file:
	cp config.example.yml config.yml
	go run main.go hello
	rm -rf config.yml

test-env:
	EXAMPLE_MESSAGE=hello_from_env_variable go run main.go hello

test-flag:
	go run main.go hello --message hello_from_cmd_flag
