package main

import (
	"log"

	"github.com/gin-gonic/gin"
	
	h "app/handlers"
)

func main() {
	r := gin.Default()

	r.POST("/signup", h.Signup)
	r.POST("/login", h.Login)
	r.POST("/changepassword", h.ChangePassword)

	port := ":4000"

	log.Println("serving: " + port)
	r.Run(port)	
}
