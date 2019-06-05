package hackathontoolkit

import (
	"fmt"

	"github.com/go-chat-bot/bot"
)

func setpizzatime(command *bot.Cmd) (msg string, err error) {
	msg = fmt.Sprintf("Reminding you to get pizza at %s", command.Args[0])
	return
}

func init() {
	bot.RegisterCommand(
		"settizzatime",
		"Tells you when the next pizza delivery is comming",
		"",
		setpizzatime)

}
