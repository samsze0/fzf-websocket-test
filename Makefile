server:
	FZF_API_KEY="test" FZF_LOG_FILE_PATH="./log.txt" $(FZF_DIR)/bin/fzf --websocket-listen="localhost:12010" --bind="start:websocket-broadcast@Hi from server@"