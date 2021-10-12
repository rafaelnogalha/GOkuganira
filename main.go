package main

import (
	"GOkuganira/controllers"
	"GOkuganira/utils"
	"fmt"
	"net"
	"net/http"

	static "github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
	"gopkg.in/olahol/melody.v1"
)

func main() {
	// port := os.Getenv("PORT");
	r := gin.Default()
	m := melody.New()
	db := utils.SetupModels() //new database

	r.Use(static.Serve("/", static.LocalFile("./server/public", true)))

	r.Use(func(c *gin.Context) {
		c.Set("db", db)
		c.Next()
	})

	// TODO: DM (/direct) /ws path is the same of main chat, so messages are bleeding into
	// each other.

	// Get main chat
	r.GET("/ws", func(c *gin.Context) {
		m.HandleRequest(c.Writer, c.Request)
	})

	// Get other room
	r.GET("/direct", func(c *gin.Context) {
		fmt.Println("GET NO DIRECT")
		http.ServeFile(c.Writer, c.Request, "./server/public/direct.html")
	})

	r.GET("/direct/ws", func(c *gin.Context) {
		fmt.Println("GET NO DIRECT WS")
		m.HandleRequest(c.Writer, c.Request)
	})

	// Get all users at the database
	r.GET("/users", controllers.FindUsers)

	// Post one user at database
	r.POST("/users", controllers.CreateUser)

	// TODO: messages are bleeding to other chatrooms, gotta fix dat
	m.HandleMessage(func(s *melody.Session, msg []byte) {
		m.BroadcastFilter(msg, func(q *melody.Session) bool {
			cond := q.Request.URL.Path == s.Request.URL.Path
			// fmt.Println("MSG ", string (msg))
			fmt.Println("PATH ", q.Request.URL.Path)
			fmt.Println("PATH ", s.Request.URL.Path)
			// fmt.Println("TRUE? ", cond)
			return cond
		})
	})

	m.HandleConnect(func(s *melody.Session) {
		fmt.Println("Number of active sessions: ", m.Len())
		ip := GetLocalIP()
		msg := `{
			"kind": "message",
			"username":"` + ip + `",
			"content": "Resolveu estragar a conversa!"
		}`

		m.BroadcastOthers([]byte(msg), s)
	})
	
	m.HandleDisconnect(func(s *melody.Session){
		fmt.Println("Number of active sessions: ", m.Len())
	})

	r.Run(":5000")
}

func GetLocalIP() string {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return ""
	}
	for _, address := range addrs {
		// check the address type and if it is not a loopback the display it
		if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				return ipnet.IP.String()
			}
		}
	}
	return ""
}
