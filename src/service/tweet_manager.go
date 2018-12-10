package service

import (
	"fmt"
	"os"
	"strings"

	"github.com/franferrari/Twitter/src/domain"
)

func main() {}

type TweetManager struct {
	tweet         domain.Tweet
	listAllTweets []domain.Tweet
	tweetsByUser  map[string][]domain.Tweet
	listAllUsers  []*domain.User
	TweetWriter   TweetWriter
}

type TweetWriter interface {
	Write(tweet domain.Tweet)
}

type MemoryTweetWriter struct {
	lastTweet domain.Tweet
}

type FileTweetWriter struct {
	file *os.File
}

func NewMemoryTweetWriter() *MemoryTweetWriter {
	return &MemoryTweetWriter{}
}

func NewFileTweetWriter() *FileTweetWriter {
	file, _ := os.OpenFile(
		"tweets.txt",
		os.O_WRONLY|os.O_TRUNC|os.O_CREATE,
		0666,
	)
	writer := new(FileTweetWriter)
	writer.file = file
	return writer
}

func (memTweetWriter *MemoryTweetWriter) Write(tweet domain.Tweet) {
	memTweetWriter.lastTweet = tweet
}

func (fileTweetWriter *FileTweetWriter) Write(tweet domain.Tweet) {
	go func() {
		if fileTweetWriter.file != nil {
			byteSlice := []byte(tweet.GetPrintableTweet() + "\n")
			fileTweetWriter.file.Write(byteSlice)
		}
	}()
}

func (memTweetWriter *MemoryTweetWriter) GetLastSavedTweet() domain.Tweet {
	return memTweetWriter.lastTweet
}

func NewTweetManager() *TweetManager {
	var tweet domain.Tweet
	listAllTweets := make([]domain.Tweet, 0)
	tweetsByUser := make(map[string][]domain.Tweet)
	listAllUsers := make([]*domain.User, 0)
	tweetWriter := NewFileTweetWriter()
	return &TweetManager{tweet, listAllTweets, tweetsByUser, listAllUsers, tweetWriter}
}

//PublishTweet publica un tweet
func (tweetMgr *TweetManager) PublishTweet(tweet domain.Tweet) (int, error) {
	if tweet.GetUser() == "" {
		return 0, fmt.Errorf("User is required")
	}
	if tweet.GetText() == "" {
		return 0, fmt.Errorf("Text is required")
	}
	if len(tweet.GetText()) >= 140 {
		return 0, fmt.Errorf("Tweet can't exceed 140 characters")
	}
	if getType(tweet) == "Image" {
		if tweet.GetImageURL() == "" {
			return 0, fmt.Errorf("Image is required")
		}
	}
	if getType(tweet) == "Quote" {
		if tweet.GetQuotedTweet() == nil {
			return 0, fmt.Errorf("Quoted tweet is required")
		}
	}
	tweetMgr.tweet = tweet
	tweetMgr.tweet.SetId(len(tweetMgr.listAllTweets) + 1)
	tweetMgr.listAllTweets = append(tweetMgr.listAllTweets, tweetMgr.tweet)
	tweetMgr.tweetsByUser[tweetMgr.tweet.GetUser()] = append(tweetMgr.tweetsByUser[tweetMgr.tweet.GetUser()], tweetMgr.tweet)
	tweetMgr.TweetWriter.Write(tweetMgr.tweet)
	return tweetMgr.tweet.GetId(), nil
}

//CountTweetsByUser devuelve la cantidad de tweets hechos por un usuario
func (tweetMgr *TweetManager) CountTweetsByUser(user string) int {
	return len(tweetMgr.tweetsByUser[user])
}

//GetTweetsByUser devuelve los tweets de un usuario
func (tweetMgr *TweetManager) GetTweetsByUser(user string) []domain.Tweet {
	_, exists := tweetMgr.tweetsByUser[user]
	if exists {
		return tweetMgr.tweetsByUser[user]
	}
	return nil
}

//GetTweet obtiene un tweet
func (tweetMgr *TweetManager) GetTweet() domain.Tweet {
	return tweetMgr.tweet
}

//GetTweets obtiene todos los tweets
func (tweetMgr *TweetManager) GetTweets() []domain.Tweet {
	return tweetMgr.listAllTweets
}

//GetTweetByID obtiene el tweet que tenga un cierto ID
func (tweetMgr *TweetManager) GetTweetByID(id int) (domain.Tweet, error) {
	var tweet domain.Tweet
	for _, tweetie := range tweetMgr.listAllTweets {
		if tweetie != nil && tweetie.GetId() == id {
			tweet = tweetie
		}
	}
	if tweet == nil {
		return nil, fmt.Errorf("Couldn't find tweet with that ID")
	}
	return tweet, nil
}

func (tweetMgr *TweetManager) SearchTweetsContaining(query string, c chan domain.Tweet) {
	go func() {
		for _, tweet := range tweetMgr.listAllTweets {
			if strings.Contains(strings.ToLower(tweet.GetText()), strings.ToLower(query)) {
				c <- tweet
			}
		}
		close(c)
	}()
}

//RegisterUser registra un usuario y lo añade a su lista de usuarios
func (tweetMgr *TweetManager) RegisterUser(user *domain.User) error {
	if user.Name == "" {
		return fmt.Errorf("Name is required")
	}
	if user.Email == "" {
		return fmt.Errorf("Email is required")
	}
	if user.Nick == "" {
		return fmt.Errorf("Nickname is required")
	}
	if user.Password == "" {
		return fmt.Errorf("Password is required")
	}
	tweetMgr.listAllUsers = append(tweetMgr.listAllUsers, user)
	return nil
}

//GetUser retorna el usuario buscado, o nil si no se encuentra en la lista de usuarios
func (tweetMgr *TweetManager) GetUser(nick string) *domain.User {
	for _, user := range tweetMgr.listAllUsers {
		if user.Nick == nick {
			return user
		}
	}
	return nil
}

func getType(tweet domain.Tweet) string {
	switch tweet.(type) {
	case *domain.TextTweet:
		return "Text"
	case *domain.ImageTweet:
		return "Image"
	case *domain.QuoteTweet:
		return "Quote"
	default:
		return "None"
	}
}
