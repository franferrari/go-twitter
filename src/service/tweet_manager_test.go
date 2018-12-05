package service_test

import (
	"testing"

	"github.com/franferrari/Twitter/src/domain"
	"github.com/franferrari/Twitter/src/service"
)

func TestPublishedTweetIsSaved(t *testing.T) {
	user := "grupoesfera"
	text := "This is my first tweet"
	tweet := domain.NewTweet(user, text)

	service.PublishTweet(tweet)

	publishedTweet := service.GetTweet()
	if publishedTweet.User != user && publishedTweet.Text != text {
		t.Errorf("Expected tweet is %s: %s \n But it's %s: %s", user, tweet, publishedTweet.User, publishedTweet.Text)
	}

	if publishedTweet.Date == nil {
		t.Error("Expected tweet can't be nil")
	}
}
