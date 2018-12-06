package main

import (
	"time"

	"github.com/abiosoft/ishell"
	"github.com/franferrari/Twitter/src/domain"
	"github.com/franferrari/Twitter/src/service"
)

func main() {
	service.InitializeService()

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

			service.PublishTweet(domain.NewTweet("Usuario", tweet))

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

	shell.AddCmd(&ishell.Cmd{
		Name: "showAllTweets",
		Help: "Shows all tweets",
		Func: func(c *ishell.Context) {

			defer c.ShowPrompt(true)

			tweets := service.GetTweets()
			if len(tweets) == 0 {
				c.Println("No existen tweets")
				return
			}

			for _, tweet := range tweets {
				if tweet != nil {
					c.Printf("User: %s\n Text: %s\n Date and time: %v\n\n", tweet.User, tweet.Text, tweet.Date.Format(time.RFC822))
				}
			}
			return
		},
	})

	shell.AddCmd(&ishell.Cmd{
		Name: "getTweetsByUser",
		Help: "Shows all tweets by a single user",
		Func: func(c *ishell.Context) {

			defer c.ShowPrompt(true)

			c.Print("Write the username: ")

			user := c.ReadLine()

			tweets := service.GetTweetsByUser(user)

			if len(tweets) == 0 {
				c.Println("No tweets for that user")
				return
			}

			for _, tweet := range tweets {
				if tweet != nil {
					c.Printf("User: %s\n Text: %s\n Date and time: %v\n\n", tweet.User, tweet.Text, tweet.Date.Format(time.RFC822))
				}
			}
			return
		},
	})

	shell.Run()
}
