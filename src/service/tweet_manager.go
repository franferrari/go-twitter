package service

import (
	"fmt"

	"github.com/franferrari/Twitter/src/domain"
)

//Tweet es el string con el comentario
var Tweet *domain.Tweet
var listAllTweets []*domain.Tweet
var tweetsByUser map[string][]*domain.Tweet

func main() {}

//InitializeService limpia la lista de tweets guardados y inicializa el mapa con los tweets de los usuarios
func InitializeService() {
	listAllTweets = make([]*domain.Tweet, 0)
	tweetsByUser = make(map[string][]*domain.Tweet)
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
	Tweet.Id = len(listAllTweets) + 1
	listAllTweets = append(listAllTweets, Tweet)
	tweetsByUser[Tweet.User] = append(tweetsByUser[Tweet.User], Tweet)
	return Tweet.Id, nil
}

//CountTweetsByUser devuelve la cantidad de tweets hechos por un usuario
func CountTweetsByUser(user string) int {
	return len(tweetsByUser[user])
}

//GetTweetsByUser devuelve los tweets de un usuario
func GetTweetsByUser(user string) []*domain.Tweet {
	_, exists := tweetsByUser[user]
	if exists {
		return tweetsByUser[user]
	}
	return nil
}

//GetTweet obtiene un tweet
func GetTweet() *domain.Tweet {
	return Tweet
}

//GetTweets obtiene todos los tweets
func GetTweets() []*domain.Tweet {
	return listAllTweets
}

//GetTweetByID obtiene el tweet que tenga un cierto ID
func GetTweetByID(id int) (*domain.Tweet, error) {
	var tweet *domain.Tweet
	for _, tweetie := range listAllTweets {
		if tweetie != nil && tweetie.Id == id {
			tweet = tweetie
		}
	}
	if tweet == nil {
		return nil, fmt.Errorf("Couldn't find tweet with that ID")
	}
	return tweet, nil
}
