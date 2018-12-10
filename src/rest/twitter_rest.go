package rest

import (
	"net/http"
	"strconv"

	"github.com/franferrari/Twitter/src/domain"
	"github.com/franferrari/Twitter/src/service"
	"github.com/gin-gonic/gin"
)

var tweetManagerServer *service.TweetManager

func NewGinServer(tweetManager *service.TweetManager) {
	tweetManagerServer = tweetManager
	router := gin.Default()
	router.GET("/tweet/:tweetId", funcionQueHaceGet)
	router.GET("/all", funcionQueHaceGetParaTodosLosTweets)
	router.GET("/search/:query", buscarPorQuery)
	router.POST("/newUser", registrarUsuario)
	router.POST("/publishTweet", publishTweet)
	router.POST("/publishImageTweet", publishImageTweet)
	router.POST("/publishQuoteTweet", publishQuoteTweet)
	go router.Run()
}

func funcionQueHaceGet(c *gin.Context) {
	tweetID, _ := strconv.Atoi(c.Param("tweetId"))
	tweet, e := tweetManagerServer.GetTweetByID(tweetID)
	if e == nil {
		c.String(http.StatusOK, tweet.GetPrintableTweet())
		c.String(200, "\n")
		c.JSON(200, gin.H{
			"tweet": tweet,
		})
	} else {
		c.String(http.StatusOK, e.Error())
	}
}

func funcionQueHaceGetParaTodosLosTweets(c *gin.Context) {
	listaTweets := tweetManagerServer.GetTweets()
	if len(listaTweets) == 0 {
		c.String(http.StatusOK, "No hay tweets!")
	} else {
		for _, tweet := range listaTweets {
			c.String(http.StatusOK, tweet.GetPrintableTweet())
			c.String(200, "\n")
			c.JSON(200, gin.H{
				"tweet": tweet,
			})
			c.String(200, "\n")
		}
	}
}

func buscarPorQuery(c *gin.Context) {
	query := c.Param("tweetId")
	searchResult := make(chan domain.Tweet)
	tweetManagerServer.SearchTweetsContaining(query, searchResult)

	foundTweets := make([]domain.Tweet, 0)

	foundTweets = append(foundTweets, <-searchResult)

	if len(foundTweets) == 0 {
		c.String(200, "No tweets were found with that string")
		return
	}

	for _, tweet := range foundTweets {
		if tweet != nil {
			c.String(http.StatusOK, tweet.GetPrintableTweet())
			c.String(200, "\n")
			c.JSON(200, gin.H{
				"tweet": tweet,
			})
			c.String(200, "\n")
		}
	}
	return
}

func registrarUsuario(c *gin.Context) {
	var usuario *domain.User
	if err := c.ShouldBindJSON(&usuario); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	//tweetManagerServer.RegisterUser()
}

func publishTweet(c *gin.Context) {
	var tweet domain.TextTweet
	if err := c.ShouldBindJSON(&tweet); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	newTweet := domain.NewTextTweet(tweet.GetUser(), tweet.GetText())

	tweetManagerServer.PublishTweet(newTweet)
}

func publishImageTweet(c *gin.Context) {
	var tweet domain.ImageTweet
	if err := c.ShouldBindJSON(&tweet); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	newTweet := domain.NewImageTweet(tweet.GetUser(), tweet.GetText(), tweet.GetImageURL())

	tweetManagerServer.PublishTweet(newTweet)
}

func publishQuoteTweet(c *gin.Context) {
	var tweet domain.ImageTweet
	if err := c.ShouldBindJSON(&tweet); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	newTweet := domain.NewQuoteTweet(tweet.GetUser(), tweet.GetText(), tweet.GetQuotedTweet())

	tweetManagerServer.PublishTweet(newTweet)
}
