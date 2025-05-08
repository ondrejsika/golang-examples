run-all:
	@make run-env
	@echo
	@make run-flag
	@echo
	@make run-no-token

run-env:
	TOKEN=from_env go run .

run-flag:
	go run . --token from_flag

run-no-token:
	go run .
