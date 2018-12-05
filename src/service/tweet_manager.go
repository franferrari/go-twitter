package service

import "github.com/franferrari/Twitter/src/domain"

//Tweet es el string con el comentario
var Tweet *domain.Tweet

func main() {

}

//PublishTweet publica un tweet
func PublishTweet(tweet *domain.Tweet) {
	Tweet = tweet
}

//GetTweet obtiene un tweet
func GetTweet() *domain.Tweet {
	return Tweet
}
