package main

// My first project while passing the learning stage
// Contains important tips about starting a project

// In the vscode terminal at the beginning of each project  Enter 'go mod init' command
// This means that we do not need to use the import section manually and this is done automatically.
// To enter project specifications, package names, etc.
// For example
// 'Go mod init MyTiwwter'
// In this case vscode a file (Go mod) creates which we enter and code the project in this file.

// The 'Go get' command is used to install the library in the terminal, libraries from the main GO site or from github.com
// The use of the variable must always occur in the Main space(expected declaration error)
// Anything other than the definition of a variable involves use, and only the definition of a variable can be done outside of this space (main).
// Variables must always be defined before use.
// The defined variable should always be used (it can be printed to avoid errors).

// The gosum file contains more information about downloading the packages used in the project.

// Using the Twigo library and the scribble package,
// we want to capture the identities of ID people on Twitter and save them Tweets on Scribble.

// Twigo from github , written by user arshamalh
// scribble is a Simple JSON Database in Golang

// To use Tiwgo, we need items such as consumer key, access token, bearer token, which we need to create.

// Twigo is a fast and easy to use twitter API library help we write best twitter bots.
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

// The words that are automatically written before the import packages are called alias,
// which is a substitute for the original name used.

// twigo.NewClient((&twigo.Config{
// ConsumerKey:    "ConsumerKey",
//   ConsumerSecret: "ConsumerSecret",
//   AccessToken:    "AccessToken",
//   AccessSecret:   "AccessTokenSecret",
// 	BearerToken:    "BearerToken",
//   }))

// The above commented code will be used later in the main section.

// Both of "bearer token" or "four other keys (ConsumerKey, ...)" is not mandatory.
// We'll find bearer token automatically if it's not specified.
// Also, you can use twigo.utils.BearerFinder() function.

// For example, we want to receive a tweet like this.
//   response, _ := client.GetTweet(tweet_id, nil)
//   fmt.Printf("%#v\n", response)
//   tweet := response.Data

// Other features :
// Making a tweet
// Retweeting and Liking a tweet
// deleting your like and retweet
// At this point, after receiving the tweets, we intend to save them in scribble.
// We need to create a client using the bearer token. This is done according to the instructions given.
// bearer token is for individuals and privately, so we do not use it in the code.
// bearer token must be through the package dotenv Load.

// Bearer token is a confidential code.
// With this code, you can access users' accounts and write robots instead.

// what is the Token in programming ?
//   A token used to log users into computer services .Security tokens are used to prove the identity of each person.
//  In terms of computer security, a "token" is a type of encrypted data in which a string of generated
// we must say that today tokens are used by users instead of using a normal password to log in to the system.
// Some tokens are encrypted keys that store digital or biometric information such as fingerprints.

// consumer key :
//  An API key (known as a consumer key) is a string value passed by a client app to your API proxies.
// The key uniquely identifies the client app.
//  access token :
// using token for User login base token authentication methods.

// In this section, we encountered an error that we fixed. Error text:
// could not import github.com/joho/godotenv
// This means that this library was not installed.
// To solve the problem we went to the terminal and ordered
// go get github.com/joho/godotenv
// And this library was installed for us.

func main() {
	godotenv.Load()
	// tweet_ID is used to get the user ID, which is a variable and needs to be defined.
	// for examlpe :
	// tweet_ID := "1598986"
	client, err := twigo.NewClient(&twigo.Config{
		// bearer token is only used to read tweets and for example we can not do activities such as liking with it.
		// We leave the following four items that we do not need empty.
		ConsumerKey:    "",
		ConsumerSecret: "",
		AccessToken:    "",
		AccessSecret:   "",
		BearerToken:    os.Getenv("BEARER_TOKEN"),
	})

	// The reason for writing the if statement in the bottom line is that in the field
	// client, err: = twigo.NewClient (& twigo.Config
	// err had an error, and typing this command, which says print if an error occurs, fixes the error.

	if err != nil {
		fmt.Println(err)
	}

	// Create the .env file to save bearer token.
	// The string in this file is our token.

	// .Env files store environmental variables.
	// an environment variable is a variable whose value is set outside the program.
	// And stores a series of information in our operating system (where it has high security)
	// We store sensitive information such as tokens in these files and never place them in open source locations such as github.
	// The godotenv.Load () command takes the .env file information and sends it to the operating system (secure location)
	// In this case, after loading this information with the said command, we must call this information using the os package,
	//  which is the operating system package.

	userResponse, err := client.GetUserByUsername("arshamalh", nil)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(userResponse.Data.ID)

	// in part of
	// client.GetUserByUsername ("arshamalh")
	// According to what the package says, we should also give those params as arguments,
	// but since we do not need it now, we leave it blank and add nill.

	db, _ := scribble.New("tweets", nil)
	user_id := userResponse.Data.ID
	tweetsResponse, _ := client.GetUserTweets(user_id, nil)
	// Since we do not need an error in this section
	// tweetsResponse, _
	// We used _.
	tweets := tweetsResponse.Data
	for _, tweet := range tweets {

		// "Tweets" is an array of tweets that must be indexed when using range to avoid errors.

		// If we want to run the program somewhere and it has
		//  an error and we do not notice the error, we use error handling.
		// if err != nil {
		// 	fmt.Println(err)
		// }

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

// We can create a file
// .gitignore
// Do something that github does not want the files in this file (folder) from us. Then the file
// .env
// We write the bearer token in this file and the second line.

// @arshamalh
// Is a username

// Tiwitter id
// 1254655 ………. Live in chat
// With this code, you can receive personal tweets and save them with scribble.

// note :
// in this line
// var e = 'f'
// The value that returns is a number because the letter and' 'is used to return asci values.

// function calls should be inside of the functions.
// Structures are defined above the mine and the functions below it
