package main

import (

	// "log"

	"fmt"
	"strings"
	"time"

	// twitch "github.com/gempir/go-twitch-irc/v3"

	"github.com/gempir/go-twitch-irc/v3"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

type PogMessage struct {
	User    string    `json:"user"`
	Message string    `json:"message"`
	Time    time.Time `json:"time"`
}

type EmoteCount struct {
	name   string
	Emotes []twitch.PrivateMessage
	count  int
}

func (e *EmoteCount) getEmoteCount() int {
	for i, emote := range e.Emotes {
		if time.Now().Sub(emote.Time).Seconds() > 60 {
			e.Emotes = append(e.Emotes[:i], e.Emotes[i+1:]...)
		}
	}
	return len(e.Emotes)
}

func (e *EmoteCount) addEmote(message twitch.PrivateMessage) []twitch.PrivateMessage {
	e.Emotes = append(e.Emotes, message)
	return e.Emotes
}

var (
	Channel     string = "swolenesss" // channel that messages will be written to
	channelList        = []string{}
	pogList            = []PogMessage{}
	PogCount    *int
	pogCount    int

	upgrader = websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
	}
)

func main() {
	chat := make(chan string, 100)

	client := twitch.NewAnonymousClient()
	go connectClient(client)

	r := gin.Default()

	// use this to serve vue app
	// r.Use(static.Serve("/", static.LocalFile("./frontend/dist", false)))
	// r.NoRoute(func(c *gin.Context) {
	// 	c.File("./frontend/dist/index.html")
	// })

	// use this to serve websocket
	r.LoadHTMLFiles("./cmd/server/index.html")
	r.GET("/", homeHandler)
	r.GET("/ws", func(c *gin.Context) {
		conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
		if err != nil {
			fmt.Printf("Failed to set websocket upgrade: %+v", err)
			return
		}

		// write messages
		go func() {
			for {
				msg := <-chat
				t := 1
				conn.WriteMessage(t, []byte(fmt.Sprintf("%s", msg)))
			}
		}()

		// signal := make(chan int, 1)
		//read messages
		// go func() {
		// 	_, signalMsg, err := conn.ReadMessage()
		// 	if err != nil {
		// 		fmt.Println("error", err)
		// 		conn.Close()
		// 		signal <- 1
		// 	}
		// 	if string(signalMsg) == "CLOSE" {
		// 		signal <- 1
		// 	}
		// }()
	})

	r.GET("/api/:name", func(c *gin.Context) {
		channel := c.Param("name")
		fmt.Printf("Channel: %s\n", channel)

		joinChannel(channel, client)

		client.OnPrivateMessage(func(message twitch.PrivateMessage) {
			if message.Channel == channel {
				// fmt.Printf("%v %v - %v\n", message.Time, message.Channel, message.Message)
				chat <- fmt.Sprintf("%v:%v - %v", message.Channel, message.User.DisplayName, message.Message)
			}
		})

		c.HTML(200, "index.html", nil)
	})

	r.GET("/api/:name/:emote", func(c *gin.Context) {

		channel := c.Param("name")
		emote := c.Param("emote")
		emoteCount := EmoteCount{name: emote}
		fmt.Printf("Channel: %s\n", channel)
		fmt.Printf("Emote: %s\n", emote)

		joinChannel(channel, client)

		client.OnPrivateMessage(func(message twitch.PrivateMessage) {
			if message.Channel == channel && strings.Contains(message.Message, emote) {
				emoteCount.addEmote(message)
			}
		})

		go func() {
			for {
				time.Sleep(500 * time.Millisecond)
				fmt.Printf("%v : %v\n", emoteCount.name, emoteCount.getEmoteCount())
				chat <- fmt.Sprintf("%v : %v", emoteCount.name, emoteCount.getEmoteCount())
			}
		}()

		c.HTML(200, "index.html", nil)
	})

	// start server
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

func contains(a string, list []string) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}

func joinChannel(channel string, client *twitch.Client) {
	if contains(channel, channelList) != true {
		client.Join(channel)
		fmt.Println("Channel joined")
	} else {
		fmt.Println("Channel already joined")
	}
}

// I dont think this is needed anymore
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

// func (c *twitch.Client) DepartAll(channel string) {
// 	if c.connActive.get() {
// 		c.send(fmt.Sprintf("PART #%s", channel))
// 	}

// 	c.channelsMtx.Lock()
// 	delete(c.channels, channel)
// 	c.channelUserlistMutex.Lock()
// 	delete(c.channelUserlist, channel)
// 	c.channelUserlistMutex.Unlock()
// 	c.channelsMtx.Unlock()
// }
