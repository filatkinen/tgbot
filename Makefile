BIN:="./bin/tgbot"

build:
	go build -v -o $(BIN) ./cmd/bot/

run: build
	$(BIN)