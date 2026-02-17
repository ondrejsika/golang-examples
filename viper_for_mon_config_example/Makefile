all:
	@make test-file
	@make test-env

test-file:
	go run .

test-env:
	MON_NOTIFICATIONS_EMAIL_SMTP_HOST=smtp.xxx.com \
	MON_NOTIFICATIONS_EMAIL_SMTP_PORT=1234 \
	MON_NOTIFICATIONS_TELEGRAM_TOKEN=telegram-token-env... \
	go run .
