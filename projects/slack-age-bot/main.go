package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/joho/godotenv"
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
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("error loading .env file")
	}
	slack_bot_token, ok := os.LookupEnv("SLACK_BOT_TOKEN")
	if !ok {
		fmt.Println("error: could not find SLACK_BOT_TOKEN")
	}
	slack_app_token, ok := os.LookupEnv("SLACK_APP_TOKEN")
	if !ok {
		fmt.Println("error: could not find SLACK_APP_TOKEN")
	}

	bot := slacker.NewClient(slack_bot_token, slack_app_token)

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

	err = bot.Listen(ctx)
	if err != nil {
		log.Fatal(err)
	}
}
