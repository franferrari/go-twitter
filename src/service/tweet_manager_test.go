package service_test

import (
	"testing"

	"github.com/franferrari/Twitter/src/domain"
	"github.com/franferrari/Twitter/src/service"
)

func TestPublishedTweetIsSaved(t *testing.T) {
	var tweet *domain.Tweet
	user := "grupoesfera"
	text := "This is my first tweet"
	tweet = domain.NewTweet(user, text)

	service.PublishTweet(tweet)

	publishedTweet := service.GetTweet()

	if publishedTweet.User != user && publishedTweet.Text != text {
		t.Errorf("Expected tweet is %s: %s \n But it's %s: %s", user, tweet, publishedTweet.User, publishedTweet.Text)
	}

	if publishedTweet.Date == nil {
		t.Errorf("Expected tweet can't be nil")
	}
}

func TestTweetWithoutUserIsNotPublished(t *testing.T) {
	var tweet *domain.Tweet
	var user string
	text := "This is my first tweet"
	tweet = domain.NewTweet(user, text)

	var err error
	err = service.PublishTweet(tweet)

	if err != nil && err.Error() != "User is required" {
		t.Error("Expected error is user is required")
	}
}

func TestTweetWithoutTextIsNotPublished(t *testing.T) {
	var tweet *domain.Tweet
	user := "usuario"
	var text string
	tweet = domain.NewTweet(user, text)

	var err error
	err = service.PublishTweet(tweet)

	if err != nil && err.Error() != "Text is required" {
		t.Error("Expected error is text is required")
	}
}

func TestTweetExceeding140CharactersIsNotPublished(t *testing.T) {
	var tweet *domain.Tweet
	user := "usuario"
	text := "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"
	tweet = domain.NewTweet(user, text)

	var err error
	err = service.PublishTweet(tweet)

	if err != nil && err.Error() != "Tweet can't exceed 140 characters" {
		t.Error("Expected error is Tweet can't exceed 140 character")
	}
}
