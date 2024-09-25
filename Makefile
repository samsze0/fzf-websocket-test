fzf-server:
	FZF_API_KEY="test" FZF_LOG_FILE_PATH="./log.txt" $(FZF_DIR)/bin/fzf --websocket-listen="localhost:12000" --bind="start:websocket-broadcast@Hi from server@" --bind="change:websocket-broadcast@change: {q}@" --bind="focus:websocket-broadcast@focus: {n}@"

test-client:
	go run ./fzf-websocket-test client

test-server:
	go run ./fzf-websocket-test server

log:
	tail -f log.txt