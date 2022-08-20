all:
	@make test-file
	@make test-env
	@make test-flag

test-file:
	go run main.go hello

test-env:
	EXAMPLE_MESSAGE=hello_from_env_variable go run main.go hello

test-flag:
	go run main.go hello --message hello_from_cmd_flag
