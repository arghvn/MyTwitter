package main

//Using the Twigo library and the scribble package,
//we want to capture the identities of ID people on Twitter and save them on Scribble.

//Twigo from github , written by user arshamalh
//scribble is a Simple JSON Database in Golang

//To use Tiwgo, we need items such as consumer key, access token, bearer token, which we need to create.

// Twigo is a fast and easy to use twitter API library help you write best twitter bots.
// Easily we can make a new client
// As we said, we need to import this library

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/arshamalh/twigo"
	"github.com/joho/godotenv"
	scribble "github.com/nanobox-io/golang-scribble"
)

//For example, we want to receive a tweet like this.

// Other features :

// Making a tweet
// Retweeting and Liking a tweet
// deleting your like and retweet

// At this point, after receiving the tweets, we intend to save them in scribble.
// We need to create a client using the bearer token. This is done according to the instructions given.
// bearer token is for individuals and privately, so we do not use it in the code.
// bearer token must be through the package dotenv Load.

type MyTweet struct {
	Text string `json: "text"`
	ID string `json: "id"`
	CreatedAt string `json: "created_at"`
}

func main() {
	godotenv.Load()
	client, err := twigo.NewClient(&twigo.Config{
		ConsumerKey:     "",
		ConsumerSecret:  "",
        AccessToken:     "",
		AccessSecret:    "",
		BearerToken:     os.Getenv("BEARER_TOKEN"),
	})

	if err != nil {
		fmt.Println(err)
	}

	userResponse, err := client.GetUserByUsername("arshamalh", nil)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(userResponse.Data.ID)

	db, _ := scribble.New("tweets", nil)
	user_id := userResponse.Data.ID
	tweetsResponse, _ := client.GetUserTweets(user_id, nil)
	tweets := tweetsResponse.Data
	for _, tweet := range tweets {
		// Write a fish to the database
		if err := db.Write("tweet", tweet.ID, tweet); err != nil {
			fmt.Println("Error", err)
		}
	}
	records, _ := db.ReadAll("tweet")
	for _, record := range records {
		var tweet MyTweet 
		json.Unmarshal([]byte(record), &tweet)
		fmt.Println(tweet)
	}
	// client.GetUserByUsername("@elonmusk")
}