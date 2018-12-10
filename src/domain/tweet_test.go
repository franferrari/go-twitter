package domain_test

import (
	"testing"

	"github.com/franferrari/Twitter/src/domain"
)

func TestTextTweetPrintsUserAndText(t *testing.T) {
	// Initialization
	tweet := domain.NewTextTweet("grupoesfera", "This is my tweet")
	// Operation
	text := tweet.GetPrintableTweet()
	// Validation
	expectedText := "@grupoesfera: This is my tweet"
	if text != expectedText {
		t.Errorf("The expected text is %s but was %s", expectedText, text)
	}
}

func TestCanCreateUser(t *testing.T) {
	// Initialization
	user := domain.NewUser("grupoesfera", "grupoesfera@grupoesfera.com", "grup", "12345")
	// Validation
	expectedUsername := "grupoesfera"
	expectedEmail := "grupoesfera@grupoesfera.com"
	expectedNick := "grup"
	expectedPw := "12345"

	if user.Name != expectedUsername && user.Email != expectedEmail && user.Nick != expectedNick && user.Password != expectedPw {
		t.Errorf("Expected a user but newUser returned another")
	}
}
