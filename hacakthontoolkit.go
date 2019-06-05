package hack

import (
	"fmt"

	"github.com/go-chat-bot/bot"
)

func setPizzaTime(command *bot.Cmd) (msg string, err error) {
	msg = fmt.Sprintf("Reminding you to get pizza at %s", command.Args[0])
	return
}

func init() {
	bot.RegisterCommand(
		"setPizzaTime",
		"Tells you when the next pizza delivery is comming",
		"",
		setPizzaTime)

}
