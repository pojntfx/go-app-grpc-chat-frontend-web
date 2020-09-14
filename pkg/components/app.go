package components

import (
	"context"
	"log"

	"github.com/maxence-charriere/go-app/v7/pkg/app"
	proto "github.com/pojntfx/go-app-grpc-chat-frontend-web/pkg/proto/generated"
)

type AppComponent struct {
	app.Compo
	chatMessageChan chan *proto.ChatMessage
	client          proto.ChatServiceClient
	stream          proto.ChatService_TransceiveMessagesClient
}

func NewAppComponent(client proto.ChatServiceClient) *AppComponent {
	return &AppComponent{chatMessageChan: make(chan *proto.ChatMessage), client: client}
}

func (c *AppComponent) Render() app.UI {
	chatComponent := NewChatComponent(c.chatMessageChan, c.handleOnCreateMessage)

	return app.Main().Body(
		app.Div().Class("container").Body(
			app.H1().Class("my-3").Body(
				app.Text("go-app gRPC Chat Frontend"),
			),
			chatComponent,
		),
	)
}

func (c *AppComponent) handleOnCreateMessage(message *proto.ChatMessage) {
	if err := c.stream.Send(message); err != nil {
		log.Fatal("could not send message", err)
	}
}

func (c *AppComponent) OnMount(ctx app.Context) {
	stream, err := c.client.TransceiveMessages(context.Background())
	if err != nil {
		log.Println("could not subscribe to messages", stream)
	}

	c.stream = stream

	go func() {
		for {
			message, err := c.stream.Recv()
			if err != nil {
				log.Println("could not receive messsage", err)
			}

			c.chatMessageChan <- message

			c.Update()
		}
	}()
}
