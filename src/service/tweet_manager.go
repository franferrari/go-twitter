package service

import (
	"fmt"

	"github.com/franferrari/Twitter/src/domain"
)

//Tweet es el string con el comentario
var Tweet *domain.Tweet
var listatweets []*domain.Tweet

func main() {}

//InitializeService limpia la lista de tweets guardados
func InitializeService() {
	listatweets = make([]*domain.Tweet, 0)
}

//PublishTweet publica un tweet
func PublishTweet(tweet *domain.Tweet) (int, error) {
	if tweet.User == "" {
		return 0, fmt.Errorf("User is required")
	}
	if tweet.Text == "" {
		return 0, fmt.Errorf("Text is required")
	}
	if len(tweet.Text) >= 140 {
		return 0, fmt.Errorf("Tweet can't exceed 140 characters")
	}
	Tweet = tweet
	Tweet.Id = len(listatweets) + 1
	listatweets = append(listatweets, tweet)
	return Tweet.Id, nil
}

//GetTweet obtiene un tweet
func GetTweet() *domain.Tweet {
	return Tweet
}

//GetTweets obtiene todos los tweets
func GetTweets() []*domain.Tweet {
	return listatweets
}

//GetTweetByID obtiene el tweet que tenga un cierto ID
func GetTweetByID(id int) (*domain.Tweet, error) {
	var tweet *domain.Tweet
	for i := 0; i < len(listatweets); i++ {
		if listatweets[i] != nil && listatweets[i].Id == id {
			tweet = listatweets[i]
		}
	}
	if tweet == nil {
		return nil, fmt.Errorf("Couldn't find tweet with that ID")
	}
	return tweet, nil
}
