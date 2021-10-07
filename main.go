package main

import (
	"GOkuganira/controllers"
	"GOkuganira/utils"
	"fmt"
	"net"
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

	r.GET("/ws", func(c *gin.Context) {
		m.HandleRequest(c.Writer, c.Request)
	})

	// Get all users at the database
	r.GET("/users", controllers.FindUsers)

	// Post one user at database
	r.POST("/users", controllers.CreateUser)

	m.HandleMessage(func(s *melody.Session, msg []byte) {
		m.Broadcast(msg)
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

	r.Run(":"+port)
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
