package domain

import (
	"fmt"
	"time"
)

type Tweet interface {
	GetUser() string
	GetText() string
	GetId() int
	GetImageURL() string
	GetQuotedTweet() Tweet
	GetDate() *time.Time
	SetId(int)
	GetPrintableTweet() string
}

//TextTweet define la estructura de un tweet de texto
type TextTweet struct {
	Id   int        `json:"id"`
	User string     `json:"user"`
	Text string     `json:"text"`
	Date *time.Time `json:"date"`
}

//ImageTweet define la estructura de un tweet que tiene una imagen
type ImageTweet struct {
	TextTweet
	ImageURL string `json:"imageURL"`
}

//QuoteTweet define la estructura de un tweet que tiene un quotetweet
type QuoteTweet struct {
	TextTweet
	TweetRef Tweet `json:"ref"`
}

//NewTweet crea un nuevo tweet
func NewTextTweet(user string, text string) *TextTweet {
	t := time.Now()
	return &TextTweet{0, user, text, &t}
}

//NewImageTweet crea un nuevo tweet de tipo image
func NewImageTweet(user string, text string, url string) *ImageTweet {
	t := time.Now()
	return &ImageTweet{TextTweet: TextTweet{0, user, text, &t}, ImageURL: url}
}

//NewQuoteTweet crea un nuevo tweet de tipo quote
func NewQuoteTweet(user string, text string, tweetRef Tweet) *QuoteTweet {
	t := time.Now()
	return &QuoteTweet{TextTweet: TextTweet{0, user, text, &t}, TweetRef: tweetRef}
}

//String muestra el tweet de forma linda
func (tweet *TextTweet) GetPrintableTweet() string {
	return fmt.Sprintf("@%s: %s\n", tweet.User, tweet.Text)
}

func (tweet *ImageTweet) GetPrintableTweet() string {
	return fmt.Sprintf("@%s: %s %s\n", tweet.User, tweet.Text, tweet.ImageURL)
}

func (tweet *QuoteTweet) GetPrintableTweet() string {
	return fmt.Sprintf(`@%s: %s "%s"`+"\n", tweet.User, tweet.Text, tweet.TweetRef.GetUser(), tweet.TweetRef.GetPrintableTweet())
}

//String muestra el user del tweet
func (tweet *TextTweet) GetUser() string {
	return tweet.User
}

//String muestra el texto del tweet
func (tweet *TextTweet) GetText() string {
	return tweet.Text
}

//String muestra el id del tweet
func (tweet *TextTweet) GetId() int {
	return tweet.Id
}

func (tweet *TextTweet) GetDate() *time.Time {
	return tweet.Date
}

func (tweet *TextTweet) SetId(id int) {
	tweet.Id = id
}

func (tweet *TextTweet) GetImageURL() string {
	return ""
}

func (tweet *ImageTweet) GetImageURL() string {
	return tweet.ImageURL
}

func (tweet *QuoteTweet) GetImageURL() string {
	return ""
}

func (tweet *TextTweet) GetQuotedTweet() Tweet {
	return nil
}

func (tweet *ImageTweet) GetQuotedTweet() Tweet {
	return nil
}

func (tweet *QuoteTweet) GetQuotedTweet() Tweet {
	return tweet.TweetRef
}

//User define la estructura de un usuario
type User struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Nick     string `json:"nick"`
	Password string `json:"password"`
}

//NewUser crea un nuevo tweet
func NewUser(name string, email string, nick string, pw string) *User {
	return &User{name, email, nick, pw}
}
