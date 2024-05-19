fzf-server:
	FZF_API_KEY="test" $(FZF_DIR)/bin/fzf --websocket-listen="localhost:12010" --bind="start:websocket-broadcast@Hi from server@"
fzf-test-client:
	go run fzf-websocket-test-client/main.go
fzf-client:
	FZF_API_KEY="test" $(FZF_DIR)/bin/fzf --websocket-listen-to="ws://localhost:12010" --bind="start:websocket-broadcast@Hi from fzf@"
fzf-test-server:
	go run fzf-websocket-test-server/main.go
log:
	tail -f /tmp/fzf.log
