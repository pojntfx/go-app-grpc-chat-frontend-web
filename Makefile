build:
	@GOARCH=wasm GOOS=js go build -o web/app.wasm cmd/go-app-grpc-chat-frontend-web-app/main.go
	@go build -o go-app-grpc-chat-frontend-web-server cmd/go-app-grpc-chat-frontend-web-server/main.go

run: build
	@./go-app-grpc-chat-frontend-web-server

clean: build
	@go clean
	@-rm web/app.wasm

dev:
	@nodemon -w . -e go,css --exec "pkill -9 go-app-grpc-chat-frontend-web-server; make run"