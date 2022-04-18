package bot

import (
	// "fmt"
	"log"
	"os"
	// "strings"
	"time"

	"golang.org/x/oauth2/clientcredentials"

	twitch "github.com/gempir/go-twitch-irc/v3"
	"github.com/joho/godotenv"
)

type PogMessage struct {
	User    string
	Message string
	Time    time.Time
}

var (
	oauth2Config *clientcredentials.Config
	// Channel      string = "swolenesss" // channel that messages will be written to
	// PogCount     *int

	// pogCount int
	// pogList  = []PogMessage{}

	BotAccessToken string
	BotUsername    string
)

// type PogCountModifier interface {
// 	IncrementPogCount()
// }
// type PogCount struct {
// 	PogCountTotal     int
// 	PogCountPerMinute int
// }
//
// func (p *PogCount) IncrementPogCount() {
// 	p.PogCountTotal++
// }

func init() {
	// load env file
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file. ERR: %s", err)
	}
}

// func CreateClient() {
// 	BotAccessToken := os.Getenv("BOT_ACCESS_TOKEN") // password for bot account to write chat messages, need to create this programatically
// 	BotUsername := os.Getenv("BOT_USERNAME")        // does not seem to matter

// 	client := twitch.NewClient(BotUsername, BotAccessToken)

// 	// client.OnPrivateMessage(func(message twitch.PrivateMessage) {
// 	// 	if strings.Contains(strings.ToLower(message.Message), "pog") {
// 	// 		x := PogMessage{message.User.Name, message.Message, message.Time}
// 	// 		pogList = append(pogList, x)

// 	// 		pogCount++
// 	// 		fmt.Println(pogCount)

// 	// 		client.Say(Channel, fmt.Sprintf("Pog has been said %v times", pogCount))
// 	// 	}
// 	// })

// 	// return client

// 	// client.Join("swolenesss")

// 	// err = client.Connect()
// 	// if err != nil {
// 	// 	panic(err)
// 	// }
// }
