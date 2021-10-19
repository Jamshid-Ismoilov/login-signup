package handlers

import (
	"log"
	"strconv"
	"math/rand"
	"time"
	"bytes"
	"fmt"
	"html/template"


	"app/vars"
	"app/config"
	r "app/redis"
	"app/databasepg"

	"github.com/gin-gonic/gin"
	"net/smtp"

)

func Signup (c *gin.Context) {

	var SignupBody vars.SignupBody
	err := c.BindJSON(&SignupBody)

	if err != nil {
		log.Fatalf("error in Signup binding: %v", err)
	}



	_, err = databasepg.SelectDBPost(SignupBody)

	if err == nil {

		c.JSON(405, "You have already logged in")
		return
	}

	var Receivers []string

	Receivers = append(Receivers, SignupBody.Email)

	fmt.Println(Receivers)

	auth := smtp.PlainAuth("", config.GSender, config.GPassword, config.GHost)

	password := RandomGenerator()

	t, err := template.ParseFiles("handlers/main.html")

	if err != nil {
		log.Fatalf("Error in parsing files: %v", err)
	}

	var body bytes.Buffer

	mimeHeaders := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n"

	body.Write([]byte(fmt.Sprintf("Test: \n%s\n\n", mimeHeaders)))

	t.Execute(&body, struct {
		Password string
	}{
		Password: password,
	})
	err = smtp.SendMail(config.GHost + ":" + config.GPort, auth, config.GSender, Receivers, body.Bytes())
	
	if err != nil {
		log.Fatalf("Error in sending mail: %v", err)
	}
	fmt.Println("OK")

	r.SetRedis(SignupBody.Email, password)

}

func RandomGenerator() string {
   rand.Seed(time.Now().UnixNano())

   randNum := rand.Intn(1000000)

   result := strconv.Itoa(randNum)

   return result
}