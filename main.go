package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/redis"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	store, _ := redis.NewStore(10, "tcp", "redis:6379", "", []byte("secret"))
	r.Use(sessions.Sessions("user", store))

	r.POST("login", login)

	requiredAuth := r.Group("/")
	requiredAuth.Use(checkUser)
	{
		requiredAuth.GET("/user", getUser)
	}
	r.Run(":8080")
}

const userIDKey = "userID"

func login(c *gin.Context) {
	log.Println("login")
	// TODO: login
	session := sessions.Default(c)
	session.Set(userIDKey, "example_id")
	if err := session.Save(); err != nil {
		log.Printf("session save error: %s", err)
	}
}

func checkUser(c *gin.Context) {
	session := sessions.Default(c)
	currentUserID := session.Get(userIDKey)
	log.Printf("current user id: %#v", fmt.Sprint(currentUserID))

	if currentUserID == nil {
		log.Println("redirect")
		c.Redirect(http.StatusPermanentRedirect, "/login")
		c.Abort()
	} else {
		log.Println("checked user")
	}
	c.Next()
}

func getUser(c *gin.Context) {
	session := sessions.Default(c)
	currentUserID := session.Get(userIDKey)
	c.JSON(http.StatusOK, gin.H{
		"id": currentUserID,
	})

}
