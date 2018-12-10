package service_test

import (
	"strings"
	"testing"

	"github.com/franferrari/Twitter/src/domain"
	"github.com/franferrari/Twitter/src/service"
)

func TestPublishedTweetIsSaved(t *testing.T) {
	tweetManager := service.NewTweetManager()
	var tweet domain.Tweet
	user := "grupoesfera"
	text := "This is my first tweet"
	tweet = domain.NewTextTweet(user, text)

	tweetManager.PublishTweet(tweet)

	publishedTweet := tweetManager.GetTweet()

	if publishedTweet.GetUser() != user && publishedTweet.GetText() != text {
		t.Errorf("Expected tweet is %s: %v \n But it's %s: %s\n", user, tweet, publishedTweet.GetUser(), publishedTweet.GetText())
	}

	if publishedTweet.GetDate() == nil {
		t.Errorf("Expected tweet can't be nil")
	}
}

func TestTweetWithoutUserIsNotPublished(t *testing.T) {
	tweetManager := service.NewTweetManager()
	var tweet *domain.TextTweet
	var user string
	text := "This is my first tweet"
	tweet = domain.NewTextTweet(user, text)

	var err error
	_, err = tweetManager.PublishTweet(tweet)

	if err != nil && err.Error() != "User is required" {
		t.Error("Expected error is user is required")
	}
}

func TestImageTweetWithoutImageIsNotPublished(t *testing.T) {
	tweetManager := service.NewTweetManager()
	var tweet domain.Tweet
	user := "asdasda"
	text := "This is my first tweet"
	imageURL := ""
	tweet = domain.NewImageTweet(user, text, imageURL)

	var err error
	_, err = tweetManager.PublishTweet(tweet)

	if err != nil && err.Error() != "Image is required" {
		t.Error("Expected error is image is required")
	}
}

func TestTweetWithoutTextIsNotPublished(t *testing.T) {
	tweetManager := service.NewTweetManager()
	var tweet *domain.TextTweet
	user := "usuario"
	var text string
	tweet = domain.NewTextTweet(user, text)

	var err error
	_, err = tweetManager.PublishTweet(tweet)

	if err != nil && err.Error() != "Text is required" {
		t.Error("Expected error is text is required")
	}
}

func TestTweetExceeding140CharactersIsNotPublished(t *testing.T) {
	tweetManager := service.NewTweetManager()
	var tweet *domain.TextTweet
	user := "usuario"
	text := "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"
	tweet = domain.NewTextTweet(user, text)

	var err error
	_, err = tweetManager.PublishTweet(tweet)

	if err != nil && err.Error() != "Tweet can't exceed 140 characters" {
		t.Error("Expected error is Tweet can't exceed 140 character")
	}
}

func TestPublishAndRetrieveMoreThanOneTweet(t *testing.T) {

	tweetManager := service.NewTweetManager()
	var tweet, tweet2 domain.Tweet
	user := "usuario"
	text := "aaaaa"
	tweet = domain.NewTextTweet(user, text)

	user2 := "usuario2"
	text2 := "bbbbb"
	tweet2 = domain.NewImageTweet(user2, text2, "ooqm.png")

	tweetManager.PublishTweet(tweet)

	tweetManager.PublishTweet(tweet2)

	publishedTweets := tweetManager.GetTweets()

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
	tweetManager := service.NewTweetManager()
	var tweet *domain.TextTweet
	id := 1
	user := "usuario"
	text := "holis"
	tweet = domain.NewTextTweet(user, text)

	id, _ = tweetManager.PublishTweet(tweet)

	publishedTweet, e := tweetManager.GetTweetByID(id)

	if e != nil && e.Error() != "Couldn't find tweet with that ID" {
		t.Error("Expected error is couldn't fint tweet with that ID")
	}

	if !isValidTweet(publishedTweet, id, user, text) {
		t.Errorf("Tweet is invalid")
		return
	}
}

func TestCantRetrieveTweetWithInexistentID(t *testing.T) {
	tweetManager := service.NewTweetManager()
	var tweet *domain.TextTweet
	idExpected := 48
	user := "usuario"
	text := "holis"
	tweet = domain.NewTextTweet(user, text)

	tweetManager.PublishTweet(tweet)

	publishedTweet, e := tweetManager.GetTweetByID(idExpected)

	if publishedTweet != nil && e != nil && e.Error() != "Couldn't find tweet with that ID" {
		t.Error("Expected error is couldn't fint tweet with that ID")
	}
}
func TestCanCountTheTweetsSentByAnUser(t *testing.T) {
	tweetManager := service.NewTweetManager()
	var tweet, tweet2, tweet3 *domain.TextTweet
	user := "usuario"
	anotherUser := "usuario2"

	text := "This is my first tweet"
	text2 := "This is my second tweet"

	tweet = domain.NewTextTweet(user, text)
	tweet2 = domain.NewTextTweet(anotherUser, text)
	tweet3 = domain.NewTextTweet(user, text2)

	tweetManager.PublishTweet(tweet)
	tweetManager.PublishTweet(tweet2)
	tweetManager.PublishTweet(tweet3)

	count := tweetManager.CountTweetsByUser(user)

	if count != 2 {
		t.Errorf("Expected count is 2, but it was %d", count)
	}
}

func TestCanRetrieveTheTweetsSentByAnUser(t *testing.T) {
	tweetManager := service.NewTweetManager()
	var tweet, tweet2, tweet3 *domain.TextTweet
	user := "usuario"
	anotherUser := "usuario2"

	text := "This is my first tweet"
	text2 := "This is my second tweet"

	tweet = domain.NewTextTweet(user, text)
	tweet2 = domain.NewTextTweet(anotherUser, text)
	tweet3 = domain.NewTextTweet(user, text2)

	tweetManager.PublishTweet(tweet)
	tweetManager.PublishTweet(tweet2)
	tweetManager.PublishTweet(tweet3)

	tweets := tweetManager.GetTweetsByUser(user)

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
	tweetManager := service.NewTweetManager()
	var tweet, tweet2, tweet3 *domain.TextTweet
	user := "usuario"
	anotherUser := "usuario2"
	user3 := "randomuser"

	text := "This is my first tweet"
	text2 := "This is my second tweet"

	tweet = domain.NewTextTweet(user, text)
	tweet2 = domain.NewTextTweet(anotherUser, text)
	tweet3 = domain.NewTextTweet(user, text2)

	tweetManager.PublishTweet(tweet)
	tweetManager.PublishTweet(tweet2)
	tweetManager.PublishTweet(tweet3)

	tweets := tweetManager.GetTweetsByUser(user3)

	if tweets != nil {
		t.Errorf("Expected count is 0, but it was %d", len(tweets))
	}
}

func TestRegisteredUserCanBeFound(t *testing.T) {
	tweetManager := service.NewTweetManager()

	var user *domain.User

	nombre := "usuario"
	email := "usuario@gmail.com"
	nick := "user"
	password := "12345"

	user = domain.NewUser(nombre, email, nick, password)

	tweetManager.RegisterUser(user)

	usuario, _ := tweetManager.GetUser(nombre)

	if usuario == nil {
		t.Errorf("Expected one user but found none")
	}
}

func TestImageTweetPrintsUserTextAndImageURL(t *testing.T) {
	// Initialization
	tweet := domain.NewImageTweet("grupoesfera", "This is my image",
		"http://www.grupoesfera.com.ar/common/img/grupoesfera.png")
	// Operation
	text := tweet.GetPrintableTweet()
	// Validation
	expectedText := "@grupoesfera: This is my image http://www.grupoesfera.com.ar/common/img/grupoesfera.png"
	if text != expectedText {
		t.Errorf("Expected a string but got another")
	}
}

func TestQuoteTweetPrintsUserTextAndQuotedTweet(t *testing.T) {

	quotedTweet := domain.NewTextTweet("grupoesfera", "This is my tweet")
	tweet := domain.NewQuoteTweet("nick", "Awesome", quotedTweet)

	text := tweet.GetPrintableTweet()
	expectedText := `@nick: Awesome "@grupoesfera: This is my tweet"`
	if text != expectedText {
		t.Errorf("Expected a string but got another")
	}
}

func TestPublishedTweetIsSavedToExternalResource(t *testing.T) {

}

func TestCanSearchForTweetContainingText(t *testing.T) {
	tweetManager := service.NewTweetManager()

	user := "usuario"
	user2 := "usuario2"
	user3 := "randomuser"

	text := "This is my first tweet"
	text2 := "This is my second tweet"

	tweet := domain.NewTextTweet(user, text)
	tweet2 := domain.NewTextTweet(user2, text)
	tweet3 := domain.NewTextTweet(user3, text2)

	tweetManager.PublishTweet(tweet)
	tweetManager.PublishTweet(tweet2)
	tweetManager.PublishTweet(tweet3)

	searchResult := make(chan domain.Tweet)
	query := "first"
	tweetManager.SearchTweetsContaining(query, searchResult)

	foundTweet := <-searchResult

	if foundTweet == nil {
		t.Errorf("Expected a tweet but got none")
	}

	if !strings.Contains(strings.ToLower(foundTweet.GetText()), strings.ToLower(query)) {
		t.Errorf("Expected the tweet to contain %v, but didn't", query)
	}
}

func BenchmarkPublishTweetWithFileTweetWriter(b *testing.B) {
	tweetManager := service.NewTweetManager()
	tweet := domain.NewTextTweet("user", "text")

	for n := 0; n < b.N; n++ {
		tweetManager.PublishTweet(tweet)
	}
}

func isValidTweet(tweet domain.Tweet, id int, user string, text string) bool {
	if tweet.GetId() != id {
		return false
	}
	if tweet.GetUser() != user {
		return false
	}
	if tweet.GetText() != text {
		return false
	}
	return true
}
