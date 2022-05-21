package main

import (

	// "log"

	// "strings"
	"fmt"
	"time"

	// twitch "github.com/gempir/go-twitch-irc/v3"

	"github.com/gempir/go-twitch-irc/v3"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

type PogMessage struct {
	User    string
	Message string
	Time    time.Time
}

var (
	Channel  string = "swolenesss" // channel that messages will be written to
	PogCount *int

	pogCount int
	pogList  = []PogMessage{}

	upgrader = websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
	}
)

func main() {
	chat := make(chan string)

	client := twitch.NewAnonymousClient()
	client.OnPrivateMessage(func(message twitch.PrivateMessage) {
		fmt.Println(message.Message)
		chat <- message.Message
	})

	r := gin.Default()
	r.LoadHTMLFiles("./cmd/server/index.html")
	r.GET("/", homeHandler)
	r.GET("/ws", func(c *gin.Context) {
		conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
		if err != nil {
			fmt.Printf("Failed to set websocket upgrade: %+v", err)
			return
		}

		for {
			msg := <-chat
			t := 1
			conn.WriteMessage(t, []byte(fmt.Sprintf("%s", msg)))
			time.Sleep(time.Second)
		}
	})

	r.GET("/twitch/:name", func(c *gin.Context) {
		fmt.Printf("Channel: %s\n", c.Param("name"))

		client.Join(c.Param("name"))
		go connectClient(client)

		c.HTML(200, "index.html", nil)

	})

	r.Run(":8080")
}

func homeHandler(c *gin.Context) {
	c.HTML(200, "index.html", nil)
}

func connectClient(client *twitch.Client) {
	err := client.Connect()
	if err != nil {
		panic(err)
	}
}

// func ReadChat(client *twitch.Client, conn *websocket.Conn) {
// 	client.OnPrivateMessage(func(message twitch.PrivateMessage) {
// 		if strings.Contains(strings.ToLower(message.Message), "pog") {
// 			x := PogMessage{message.User.Name, message.Message, message.Time}
// 			pogList = append(pogList, x)

// 			pogCount++
// 			fmt.Println(pogCount)

// 			client.Say(Channel, fmt.Sprintf("Pog has been said %v times", pogCount))
// 			t := 1
// 			conn.WriteMessage(t, []byte(fmt.Sprintf("sup %s", message.Message)))
// 		}
// 	})
// }
