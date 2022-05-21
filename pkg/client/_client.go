// package main

// import (
// 	"context"
// 	"fmt"
// 	"log"
// 	"net/http"
// 	"os"

// 	"golang.org/x/oauth2/clientcredentials"
// 	"golang.org/x/oauth2/twitch"

// 	"github.com/joho/godotenv"
// )

// var (
// 	clientID string
// 	// Consider storing the secret in an environment variable or a dedicated storage system.
// 	clientSecret string
// 	oauth2Config *clientcredentials.Config
// )

// func main() {
// 	err := godotenv.Load("../.env")
// 	if err != nil {
// 		log.Fatalf("Error loading .env file. ERR: %s", err)
// 	}

// 	clientID := os.Getenv("CLIENT_ID")
// 	clientSecret := os.Getenv("CLIENT_SECRET")

// 	oauth2Config = &clientcredentials.Config{
// 		ClientID:     clientID,
// 		ClientSecret: clientSecret,
// 		TokenURL:     twitch.Endpoint.TokenURL,
// 	}

// 	token, err := oauth2Config.Token(context.Background())
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	fmt.Printf("Token: %s\n", token)
// 	fmt.Printf("Access token: %s\n", token.AccessToken)
// 	fmt.Println("-----------------------")

// 	// client := twitch.NewClient()
// 	// url := "https://api.twitch.tv/helix/users?login=swolenesss"
// 	url := "https://id.twitch.tv/oauth2/authorize?response_type=code&" + clientID + "+channel%3Aread%3Apolls&state=c3ab8aa609ea11e793ae92361f002671"
// 	client := http.Client{}
// 	req, err := http.NewRequest("GET", url, nil)
// 	if err != nil {
// 		log.Fatalf("ERROR = %v", err)
// 	}

// 	req.Header = http.Header{
// 		"Content-Type":  []string{"application/json"},
// 		"Authorization": []string{fmt.Sprintf("Bearer %v", token.AccessToken)},
// 	}
// 	res, err := client.Do(req)
// 	if err != nil {
// 		log.Fatalf("ERROR = %v", err)
// 	}
// 	fmt.Printf("Response: %v\n", res)

// 	// response, err = http.Get("https://api.twitch.tv/helix/users?login=twitchdev",
// 	// 	http.Header{"Authorization": fmt.Sprintf("Bearer %v", token.AccessToken)})
// }
