package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/shomali11/slacker"
)

func printCommandEvents(analyticsChannel <-chan *slacker.CommandEvent) {
	for event := range analyticsChannel {
		fmt.Println("Command Events")
		fmt.Println(event.Timestamp)
		fmt.Println(event.Command)
		fmt.Println(event.Parameters)
		fmt.Println(event.Event)
		fmt.Println()
	}
}

func main() {
	os.Setenv("SLACK_BOT_TOKEN", "xoxb-3486741035872-3460098029845-DwY4f0OYMTOj4ttbW7jXvEgC")
	os.Setenv("SLACK_APP_TOKEN", "xapp-1-A03DP9Z29EG-3456351293158-38c65cf0b702b32bf19ef062bd4eb0838c7508c960417b1928581bbf23db3840")

	bot := slacker.NewClient(os.Getenv("SLACK_BOT_TOKEN"), os.Getenv("SLACK_APP_TOKEN"))

	go printCommandEvents(bot.CommandEvents())

	bot.Command("My Year of Birth is <Year>", &slacker.CommandDefinition{
		Description: "Year of Birth Calculator",
		Example:     "My Year of Birth is 2000",
		Handler: func(botCtx slacker.BotContext, request slacker.Request, response slacker.ResponseWriter) {
			year := request.Param("Year")
			yearNum, err := strconv.Atoi(year)
			if err != nil {
				fmt.Println(err)
			}
			age := time.Now().Year() - yearNum
			r := fmt.Sprintf("your age is %d", age)
			response.Reply(r)
		},
	})

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	err := bot.Listen(ctx)
	if err != nil {
		log.Fatal(err)
	}
}
