package handlers

import (
	"github.com/gin-gonic/gin"

	"log"

	"app/vars"
	"app/databasepg"
)

func ChangePassword(c *gin.Context) {

	var body vars.ChangePasswordBody

	err := c.BindJSON(&body)

	if err != nil {
		log.Fatalf("error in change password binding: %v", err)
	}

	result := databasepg.ChangePasswordDB(body)

	if result == 0 {
	 log.Println("ChangePassword db error:", err)
	 c.JSON(405, "Password or email is incorrect")	
	return
	}
	if result == 1 {

	c.JSON(200, "Passwordchanged successfully")

	}
}