package main

import (
	"time"

	"github.com/abiosoft/ishell"
	"github.com/franferrari/Twitter/src/domain"
	"github.com/franferrari/Twitter/src/service"
)

func main() {

	shell := ishell.New()
	shell.SetPrompt("Tweeter >> ")
	shell.Print("Type 'help' to know commands\n")

	shell.AddCmd(&ishell.Cmd{
		Name: "publishTweet",
		Help: "Publishes a tweet",
		Func: func(c *ishell.Context) {

			defer c.ShowPrompt(true)

			c.Print("Write your tweet: ")

			tweet := c.ReadLine()

			service.PublishTweet(domain.NewTweet("usuario", tweet))

			c.Print("Tweet sent\n")

			return
		},
	})

	shell.AddCmd(&ishell.Cmd{
		Name: "showTweet",
		Help: "Shows a tweet",
		Func: func(c *ishell.Context) {

			defer c.ShowPrompt(true)

			tweet := service.GetTweet()
			if tweet != nil {
				c.Printf("User: %s\n Text: %s\n Date and time: %v", tweet.User, tweet.Text, tweet.Date.Format(time.RFC822))
			} else {
				c.Println("No existen tweets anteriores")
			}
			return
		},
	})

	shell.Run()

}