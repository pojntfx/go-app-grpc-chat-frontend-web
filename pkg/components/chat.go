package components

import (
	"github.com/maxence-charriere/go-app/v7/pkg/app"
	proto "github.com/pojntfx/go-app-grpc-chat-frontend-web/pkg/proto/generated"
)

type ChatComponent struct {
	app.Compo
	chatMessageChan chan *proto.ChatMessage
	onCreateMessage func(*proto.ChatMessage)

	newMessageContent string
	messages          []*proto.ChatMessage
}

func NewChatComponent(chatMessageChan chan *proto.ChatMessage, onCreateMessage func(*proto.ChatMessage)) *ChatComponent {
	return &ChatComponent{chatMessageChan: chatMessageChan, onCreateMessage: onCreateMessage}
}

func (c *ChatComponent) Render() app.UI {
	return app.Div().Body(
		app.Ul().Class("list-group").Body(
			app.Range(c.messages).Slice(func(i int) app.UI {
				return app.Li().Class("list-group-item").Body(
					app.Text(c.messages[i].GetContent()),
				)
			}),
		),
		app.Div().Class("input-group my-3").Body(
			app.Input().Type("text").Class("form-control").Value(c.newMessageContent).Placeholder("New message").OnInput(c.handleOnInput).OnChange(c.handleOnChange),
			app.Div().Class("input-group-append").Body(
				app.Button().Class("btn btn-primary").Body(app.Text("Send Message")).OnClick(c.handleOnChange),
			),
		),
	)
}

func (c *ChatComponent) handleOnInput(ctx app.Context, e app.Event) {
	c.newMessageContent = e.Get("target").Get("value").String()

	c.Update()
}

func (c *ChatComponent) handleOnChange(ctx app.Context, e app.Event) {
	message := proto.ChatMessage{Content: c.newMessageContent}

	c.newMessageContent = ""
	c.Update()

	go c.onCreateMessage(&message)
}

func (c *ChatComponent) OnMount(ctx app.Context) {
	go func() {
		for message := range c.chatMessageChan {
			c.messages = append(c.messages, message)

			c.Update()
		}
	}()
}
