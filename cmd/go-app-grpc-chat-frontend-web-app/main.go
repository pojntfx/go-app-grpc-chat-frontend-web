package main

import (
	"time"

	"github.com/maxence-charriere/go-app/v7/pkg/app"
	"github.com/pojntfx/go-app-grpc-chat-frontend-web/pkg/components"
	proto "github.com/pojntfx/go-app-grpc-chat-frontend-web/pkg/proto/generated"
	"github.com/pojntfx/go-app-grpc-chat-frontend-web/pkg/websocketproxy"
	"google.golang.org/grpc"
)

func main() {
	proxy := websocketproxy.NewWebSocketProxyClient(time.Minute)

	conn, err := grpc.Dial("ws:///127.0.0.1:10000", grpc.WithContextDialer(proxy.Dialer), grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	defer conn.Close()
	chatClient := proto.NewChatServiceClient(conn)

	appComponent := components.NewAppComponent(chatClient)

	app.Route("/", appComponent)

	app.Run()
}
