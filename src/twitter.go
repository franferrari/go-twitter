package main

import (
	"strings"
	"time"

	"github.com/abiosoft/ishell"
	"github.com/franferrari/Twitter/src/domain"
	"github.com/franferrari/Twitter/src/rest"
	"github.com/franferrari/Twitter/src/service"
)

func main() {
	tweetManager := service.NewTweetManager()
	rest.NewGinServer(tweetManager)

	shell := ishell.New()
	shell.SetPrompt("Tweeter >> ")
	shell.Print("Type 'help' to know commands\n")

	shell.AddCmd(&ishell.Cmd{
		Name: "publishTweet",
		Help: "Publishes a tweet",
		Func: func(c *ishell.Context) {

			defer c.ShowPrompt(true)

			c.Print("Select your tweet type (Text / Image / Quote): ")

			tweetType := strings.ToLower(c.ReadLine())

			switch tweetType {
			case "text":
				c.Print("Write your tweet: ")
				tweet := c.ReadLine()
				tweetManager.PublishTweet(domain.NewTextTweet("Usuario", tweet))
			case "image":
				c.Print("Write your tweet: ")
				textTweet := c.ReadLine()
				c.Print("Write your image URL: ")
				imageTweet := c.ReadLine()
				tweetManager.PublishTweet(domain.NewImageTweet("Usuario", textTweet, imageTweet))
			case "quote":
				c.Print("This isn't implemented yet, Sorry! :)")
				//tweetManager.PublishTweet(domain.NewQuoteTweet("Usuario", textTweet, ))
			}
			c.Print("Tweet sent\n")
			return
		},
	})

	shell.AddCmd(&ishell.Cmd{
		Name: "showAllTweets",
		Help: "Shows all tweets",
		Func: func(c *ishell.Context) {

			defer c.ShowPrompt(true)

			tweets := tweetManager.GetTweets()
			if len(tweets) == 0 {
				c.Println("No existen tweets")
				return
			}
			c.Println(len(tweets))
			for _, tweet := range tweets {
				if tweet != nil {
					c.Printf("User: %s\n Text: %s\n Date and time: %v\n\n", tweet.GetUser(), tweet.GetText(), tweet.GetDate().Format(time.RFC822))
				}
			}
			return
		},
	})

	shell.AddCmd(&ishell.Cmd{
		Name: "showTweetsByUser",
		Help: "Shows all tweets by a single user",
		Func: func(c *ishell.Context) {

			defer c.ShowPrompt(true)

			c.Print("Write the username: ")

			user := c.ReadLine()

			tweets := tweetManager.GetTweetsByUser(user)

			if len(tweets) == 0 {
				c.Println("No tweets for that user")
				return
			}

			for _, tweet := range tweets {
				if tweet != nil {
					c.Printf("User: %s\n Text: %s\n Date and time: %v\n\n", tweet.GetUser(), tweet.GetText(), tweet.GetDate().Format(time.RFC822))
				}
			}
			return
		},
	})

	shell.AddCmd(&ishell.Cmd{
		Name: "searchTweets",
		Help: "Search for tweets containing a specific string",
		Func: func(c *ishell.Context) {

			defer c.ShowPrompt(true)

			c.Print("What do you want to search for?: ")

			query := c.ReadLine()

			searchResult := make(chan domain.Tweet)
			tweetManager.SearchTweetsContaining(query, searchResult)

			foundTweets := make([]domain.Tweet, 0)

			foundTweets = append(foundTweets, <-searchResult)

			if len(foundTweets) == 0 {
				c.Println("No tweets were found with that string")
				return
			}

			for _, tweet := range foundTweets {
				if tweet != nil {
					c.Printf("User: %s\n Text: %s\n Date and time: %v\n\n", tweet.GetUser(), tweet.GetText(), tweet.GetDate().Format(time.RFC822))
				}
			}
			return
		},
	})

	shell.Run()
}
