package domain

import "time"

//Tweet define la estructura de un tweet
type Tweet struct {
	User string
	Text string
	Date *time.Time
}

//NewTweet crea un nuevo tweet
func NewTweet(user string, text string) *Tweet {
	t := time.Now()
	return &Tweet{user, text, &t}
}
