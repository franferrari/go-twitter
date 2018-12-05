package service

import (
	"fmt"

	"github.com/franferrari/Twitter/src/domain"
)

//Tweet es el string con el comentario
var Tweet *domain.Tweet

func main() {

}

//PublishTweet publica un tweet
func PublishTweet(tweet *domain.Tweet) error {
	if tweet.User == "" {
		return fmt.Errorf("User is required")
	}
	if tweet.Text == "" {
		return fmt.Errorf("Text is required")
	}
	if len(tweet.Text) >= 140 {
		return fmt.Errorf("Tweet can't exceed 140 characters")
	}
	Tweet = tweet
	return nil
}

//GetTweet obtiene un tweet
func GetTweet() *domain.Tweet {
	return Tweet
}
