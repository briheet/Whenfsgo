.PHONY: all server client

all: server client

server: 
	@go run server/server.go

client: 
	@go run client/client.go
