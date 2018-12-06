package domain

import "time"

//Tweet define la estructura de un tweet
type Tweet struct {
	Id   int
	User string
	Text string
	Date *time.Time
}

//NewTweet crea un nuevo tweet
func NewTweet(user string, text string) *Tweet {
	t := time.Now()
	return &Tweet{0, user, text, &t}
}
