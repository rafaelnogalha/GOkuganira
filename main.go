package main

import (
	"GOkuganira/controllers"
	"GOkuganira/utils"
	"fmt"
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

	// Get main chat
	r.GET("/ws", func(c *gin.Context) {
		m.HandleRequest(c.Writer, c.Request)
	})

	// Get other room
	r.GET("/direct/:name", func(c *gin.Context) {
		fmt.Println("GET NO DIRECT")
		http.ServeFile(c.Writer, c.Request, "./server/public/direct.html")
	})

	r.GET("/direct/:name/ws", func(c *gin.Context) {
		fmt.Println("GET NO DIRECT WS")
		m.HandleRequest(c.Writer, c.Request)
	})

	// Get all users at the database
	r.GET("/users", controllers.FindUsers)

	// Post one user at database
	r.POST("/users", controllers.CreateUser)

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

	// TODO: see if can be done with username instead of IP :)
	m.HandleConnect(func(s *melody.Session) {
		// fmt.Println("Number of active sessions: ", m.Len())
		// ip := GetLocalIP()
		// msg := `{
		// 	"kind": "message",
		// 	"username":"` + ip + `",
		// 	"content": "Resolveu estragar a conversa!"
		// }`

		// m.BroadcastOthers([]byte(msg), s)
	})
	
	m.HandleDisconnect(func(s *melody.Session){
		fmt.Println("Number of active sessions: ", m.Len())
	})

	r.Run(":5000")
}
