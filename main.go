package main

import (
	"GOkuganira/controllers"
	"GOkuganira/utils"
	"fmt"
	"net/http"
	"os"

	static "github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
	"gopkg.in/olahol/melody.v1"
)

func main() {
	port := os.Getenv("PORT");
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
	r.GET("/channel/:name/", func(c *gin.Context) {
		http.ServeFile(c.Writer, c.Request, "./server/public/channel.html")
	})

	r.GET("/channel/:name/ws", func(c *gin.Context) {
		m.HandleRequest(c.Writer, c.Request)
	})

	// Get all users at the database
	r.GET("/users", controllers.FindUsers)

	// Post one user at database
	r.POST("/users", controllers.CreateUser)

	m.HandleMessage(func(s *melody.Session, msg []byte) {
		m.BroadcastFilter(msg, func(q *melody.Session) bool {
			cond := q.Request.URL.Path == s.Request.URL.Path
			return cond
		})
	})

	m.HandleConnect(func(s *melody.Session) {
	})
	
	m.HandleDisconnect(func(s *melody.Session){
		fmt.Println("Number of active sessions: ", m.Len())
	})

	r.Run(":" + port)
}
