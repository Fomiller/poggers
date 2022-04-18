package main

import (
	"fmt"
	// "log"
	"net/http"
	// "strings"
	"time"

	// client "github.com/fomiller/poggers/pkg/twitch"
	// twitch "github.com/gempir/go-twitch-irc/v3"

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
	r := gin.Default()
	r.LoadHTMLFiles("./cmd/server/index.html")
	r.GET("/", homeHandler)
	r.GET("/ws", websocketHandler)

	r.Run(":8080")
}

func homeHandler(c *gin.Context) {
	c.HTML(200, "index.html", nil)
}

func websocketHandler(c *gin.Context) {
	wsHandler(c.Writer, c.Request)
}

func wsHandler(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Printf("Failed to set websocket upgrade: %+v", err)
		return
	}
	fmt.Println(r.URL.RequestURI())
	for {
		// if err != nil {
		// 	break
		// }
		// print(t)
		msg := "hello"
		t := 1
		conn.WriteMessage(t, []byte(fmt.Sprintf("SUP %s", msg)))
		time.Sleep(time.Second)
	}
}

// func ReadChat(client *twitch.Client) {
// 	client.OnPrivateMessage(func(message twitch.PrivateMessage) {
// 		if strings.Contains(strings.ToLower(message.Message), "pog") {
// 			x := PogMessage{message.User.Name, message.Message, message.Time}
// 			pogList = append(pogList, x)

// 			pogCount++
// 			fmt.Println(pogCount)

// 			client.Say(Channel, fmt.Sprintf("Pog has been said %v times", pogCount))
// 		}
// 	})

// }
