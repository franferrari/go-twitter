package service_test

import (
	"testing"

	"github.com/franferrari/Twitter/src/domain"
	"github.com/franferrari/Twitter/src/service"
)

func TestPublishedTweetIsSaved(t *testing.T) {
	service.InitializeService()
	var tweet *domain.Tweet
	user := "grupoesfera"
	text := "This is my first tweet"
	tweet = domain.NewTweet(user, text)

	service.PublishTweet(tweet)

	publishedTweet := service.GetTweet()

	if publishedTweet.User != user && publishedTweet.Text != text {
		t.Errorf("Expected tweet is %s: %v \n But it's %s: %s\n", user, tweet, publishedTweet.User, publishedTweet.Text)
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
	_, err = service.PublishTweet(tweet)

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
	_, err = service.PublishTweet(tweet)

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
	_, err = service.PublishTweet(tweet)

	if err != nil && err.Error() != "Tweet can't exceed 140 characters" {
		t.Error("Expected error is Tweet can't exceed 140 character")
	}
}

func TestPublishAndRetrieveMoreThanOneTweet(t *testing.T) {

	service.InitializeService()
	var tweet, tweet2 *domain.Tweet
	user := "usuario"
	text := "aaaaa"
	tweet = domain.NewTweet(user, text)

	user2 := "usuario2"
	text2 := "bbbbb"
	tweet2 = domain.NewTweet(user2, text2)

	service.PublishTweet(tweet)

	service.PublishTweet(tweet2)

	publishedTweets := service.GetTweets()

	if len(publishedTweets) != 2 {
		t.Errorf("Expected size is 2 but it was %d", len(publishedTweets))
		return
	}

	firstPublishedTweet := publishedTweets[0]
	secondPublishedTweet := publishedTweets[1]

	if !isValidTweet(firstPublishedTweet, 1, user, text) {
		t.Errorf("Tweet is invalid")
		return
	}

	if !isValidTweet(secondPublishedTweet, 2, user2, text2) {
		t.Errorf("Tweet is invalid")
		return
	}
}

func TestCanRetrieveTweetById(t *testing.T) {
	service.InitializeService()
	var tweet *domain.Tweet
	id := 1
	user := "usuario"
	text := "holis"
	tweet = domain.NewTweet(user, text)

	id, _ = service.PublishTweet(tweet)

	publishedTweet, e := service.GetTweetByID(id)

	if e != nil && e.Error() != "Couldn't find tweet with that ID" {
		t.Error("Expected error is couldn't fint tweet with that ID")
	}

	if !isValidTweet(publishedTweet, id, user, text) {
		t.Errorf("Tweet is invalid")
		return
	}
}

func TestCantRetrieveTweetWithInexistentID(t *testing.T) {
	service.InitializeService()
	var tweet *domain.Tweet
	idExpected := 48
	user := "usuario"
	text := "holis"
	tweet = domain.NewTweet(user, text)

	service.PublishTweet(tweet)

	publishedTweet, e := service.GetTweetByID(idExpected)

	if publishedTweet != nil && e != nil && e.Error() != "Couldn't find tweet with that ID" {
		t.Error("Expected error is couldn't fint tweet with that ID")
	}
}
func TestCanCountTheTweetsSentByAnUser(t *testing.T) {
	service.InitializeService()
	var tweet, tweet2, tweet3 *domain.Tweet
	user := "usuario"
	anotherUser := "usuario2"

	text := "This is my first tweet"
	text2 := "This is my second tweet"

	tweet = domain.NewTweet(user, text)
	tweet2 = domain.NewTweet(anotherUser, text)
	tweet3 = domain.NewTweet(user, text2)

	service.PublishTweet(tweet)
	service.PublishTweet(tweet2)
	service.PublishTweet(tweet3)

	count := service.CountTweetsByUser(user)

	if count != 2 {
		t.Errorf("Expected count is 2, but it was %d", count)
	}
}

func TestCanRetrieveTheTweetsSentByAnUser(t *testing.T) {
	service.InitializeService()
	var tweet, tweet2, tweet3 *domain.Tweet
	user := "usuario"
	anotherUser := "usuario2"

	text := "This is my first tweet"
	text2 := "This is my second tweet"

	tweet = domain.NewTweet(user, text)
	tweet2 = domain.NewTweet(anotherUser, text)
	tweet3 = domain.NewTweet(user, text2)

	service.PublishTweet(tweet)
	service.PublishTweet(tweet2)
	service.PublishTweet(tweet3)

	tweets := service.GetTweetsByUser(user)

	if len(tweets) != 2 {
		t.Errorf("Expected count is 2, but it was %d", len(tweets))
	}

	firstPublishedTweet := tweets[0]
	secondPublishedTweet := tweets[1]

	if !isValidTweet(firstPublishedTweet, 1, user, text) {
		t.Errorf("Tweet is invalid")
		return
	}

	if !isValidTweet(secondPublishedTweet, 3, user, text2) {
		t.Errorf("Tweet is invalid")
		return
	}
}

func TestIfUserDoesntExistReturnsNilTweets(t *testing.T) {
	service.InitializeService()
	var tweet, tweet2, tweet3 *domain.Tweet
	user := "usuario"
	anotherUser := "usuario2"
	user3 := "randomuser"

	text := "This is my first tweet"
	text2 := "This is my second tweet"

	tweet = domain.NewTweet(user, text)
	tweet2 = domain.NewTweet(anotherUser, text)
	tweet3 = domain.NewTweet(user, text2)

	service.PublishTweet(tweet)
	service.PublishTweet(tweet2)
	service.PublishTweet(tweet3)

	tweets := service.GetTweetsByUser(user3)

	if tweets != nil {
		t.Errorf("Expected count is 0, but it was %d", len(tweets))
	}
}

func isValidTweet(tweet *domain.Tweet, id int, user string, text string) bool {
	if tweet.Id != id {
		return false
	}
	if tweet.User != user {
		return false
	}
	if tweet.Text != text {
		return false
	}
	return true
}
