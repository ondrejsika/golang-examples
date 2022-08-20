all:
	@make test-file
	@make test-env
	@make test-flag

test-file:
	go run main.go

test-env:
	EXAMPLE_TOKEN=token_from_env_variable go run main.go

test-flag:
	go run main.go --token token_from_cmd_flag
