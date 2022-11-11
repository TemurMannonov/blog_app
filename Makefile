swag-init:
	swag init -g api/server.go -o api/docs

start:
	go run main.go