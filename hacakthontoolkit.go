package hackathontoolkit

import (
	"fmt"
	"time"
	"github.com/go-chat-bot/bot"
)

func setpizzatime(command *bot.Cmd) (msg string, err error) {
	msg = fmt.Sprintf("Reminding you to get pizza at %s", command.Args[0])
	return
}

func usertime(command *bot.Cmd) (msg string, err error) {
	t:= time.Now()
	msg = fmt.Sprintf("%s",t)
	return
}

func init() {
	bot.RegisterCommand(
		"settizzatime",
		"Tells you when the next pizza delivery is comming",
		"",
		setpizzatime)

		bot.RegisterCommand(
			"gettime",
			"Tells you current time",
			"",
			usertime)

}
