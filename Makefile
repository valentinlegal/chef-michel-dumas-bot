include .env

APP_NAME=chef-michel-dumas-bot

# Colors
ERROR_COLOR=\033[0;31m
NO_COLOR=\033[m
OK_COLOR=\033[0;32m

# Strings
ERROR_STRING=[ERROR]
OK_STRING=[OK]

build:
	@echo "$(NO_COLOR)Build step..."
	go build -o ${APP_NAME}
	@echo "$(OK_COLOR)$(OK_STRING)$(NO_COLOR)"

install: packages tests build

packages:
	@echo "$(NO_COLOR)Packages step..."
	go mod download -x all
	@echo "$(OK_COLOR)$(OK_STRING)$(NO_COLOR)"

release:
	@echo "$(NO_COLOR)Release app..."
	rm -rf releases/app && mkdir -p releases/app/data
	docker run --rm -it -v "${PWD}":/app chef-michel-dumas-bot sh -c "make install"
	mv ${APP_NAME} releases/app/${APP_NAME}
	cp releases/install.sh releases/app/install.sh
	cp .env.template releases/app/.env.template
	cp data/data.json releases/app/data/data.json
	tar -cvzf releases/${APP_NAME}-${BOT_VERSION}.tar.gz -C releases/ app/
	@echo "$(OK_COLOR)$(OK_STRING)$(NO_COLOR)"

run:
	@echo "$(NO_COLOR)Run step..."
	go run main.go
	@echo "$(OK_COLOR)$(OK_STRING)$(NO_COLOR)"

tests:
	@echo "$(NO_COLOR)Tests step..."
	go test -cover ./...
	@echo "$(OK_COLOR)$(OK_STRING)$(NO_COLOR)"
