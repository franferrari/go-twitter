package service

//Tweet es el string con el comentario
var Tweet string

func main() {

}

//PublishTweet publica un tweet
func PublishTweet(tweet string) string {
	Tweet = tweet
	return Tweet
}
