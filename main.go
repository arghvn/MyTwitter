package main

//My first project after passing the learning stage

//Using the Twigo library and the scribble package,
//we want to capture the identities of ID people on Twitter and save them on Scribble.

//Twigo from github , written by user arshamalh
//scribble from json database
//scribble is a Simple JSON Database in Golang

//To use Tiwgo, we need items such as consumer key, access token, bearer token, which we need to create.


// Twigo is a fast and easy to use twitter API library help you write best twitter bots.
// Easily we can make a new client
// As we said, we need to import this library

import (
	 "github.com/arshamalh/twigo"
       "github.com/joho/godotenv"
)


twigo.NewClient((&twigo.Config{

  ConsumerKey:    "ConsumerKey",

	ConsumerSecret: "ConsumerSecret",

	AccessToken:    "AccessToken",

	AccessSecret:   "AccessTokenSecret",

  // Both of "bearer token" or "four other keys (ConsumerKey, ...)" is not mandatory.
  // We'll find bearer token automatically if it's not specified.
  // Also, you can use twigo.utils.BearerFinder() function.

	BearerToken:    "BearerToken",

}))

//For example, we want to receive a tweet like this.

response, _ := client.GetTweet(tweet_id, nil)

fmt.Printf("%#v\n", response)

tweet := response.Data

// Other features :
 
//Making a tweet
//Retweeting and Liking a tweet
//deleting your like and retweet

//At this point, after receiving the tweets, we intend to save them in scribble.

//We need to create a client using the bearer token. This is done according to the instructions given.

//bearer token is for individuals and privately, so we do not use it in the code.

// bearer token must be through the package dotenv Load.

func main() {
	godotenv.Load()
	client, err := twigo.NewClient((&tiwgo.Config{
		ConsumerKey:     os.Getenv("consumerkey"),
		ConsumerSecret:  os.Getenv("consumersecret"),
        AccessToken:     os.Getenv("accesstoken"),
		AccessSecret:    os.Getenv("accesstokensecret"),
		BearerToken:     os.Getenv("bearertoken"),
	}))
}