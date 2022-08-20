all:
	@make test-file
	@make test-env

test-file:
	go run main.go

test-env:
	EXAMPLE_TOKEN=token_from_env_variable go run main.go
