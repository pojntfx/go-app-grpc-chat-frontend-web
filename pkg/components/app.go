package components

import (
	"context"
	"log"

	"github.com/maxence-charriere/go-app/v7/pkg/app"
	proto "github.com/pojntfx/go-app-grpc-chat-frontend-web/pkg/proto/generated"
)

type App struct {
	app.Compo
	client            proto.ChatServiceClient
	receivedMessages  []proto.ChatMessage
	newMessageContent string
}

func NewApp(client proto.ChatServiceClient) *App {
	return &App{client: client, receivedMessages: []proto.ChatMessage{}, newMessageContent: ""}
}

func (c *App) Render() app.UI {
	return app.Main().Body(
		app.Div().Class("container").Body(
			app.H1().Class("mt-3").Body(
				app.Text("go-app gRPC Chat Frontend"),
			),
			app.U().Class("list-group mt-3").Body(
				app.Range(c.receivedMessages).Slice(func(i int) app.UI {
					return app.Li().Class("list-group-item").Body(
						app.Text(c.receivedMessages[i].GetContent()),
					)
				}),
			),
			app.Div().Class("input-group mt-3").Body(
				app.Input().Type("text").Class("form-control").Value(c.newMessageContent).Placeholder("Message content").OnInput(func(ctx app.Context, e app.Event) {
					c.newMessageContent = e.Get("target").Get("value").String()

					c.Update()
				}),
				app.Div().Class("input-group-append").Body(
					app.Button().Class("btn btn-primary").Body(app.Text("Send Message")).OnClick(func(ctx app.Context, e app.Event) {
						log.Println("Sending message with content", c.newMessageContent)

						message := proto.ChatMessage{Content: c.newMessageContent}
						outMessage, err := c.client.CreateMessage(context.TODO(), &message)
						if err != nil {
							log.Println("could not send message", err)
						}

						c.receivedMessages = append(c.receivedMessages, *outMessage)

						log.Println("Received from server message", outMessage)

						c.newMessageContent = ""

						c.Update()
					}),
				),
			),
		),
	)
}
