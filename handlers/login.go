package handlers

import (
	"log"
	"time"

	"app/vars"
	r "app/redis"
	"app/databasepg"

	"github.com/gin-gonic/gin"
	jwt "github.com/dgrijalva/jwt-go"
)

var jwtKey = []byte("app_jwt_login")

func Login (c *gin.Context) {
	var LoginBody vars.LoginBody
	err := c.BindJSON(&LoginBody)

	if err != nil {
		log.Fatalf("error in Login binding: %v", err)
	}

	pass, err := r.GetRedis(LoginBody.Email)

	if err != nil {
		log.Printf("error in getting redis data: %v", err)

	}

	if pass == LoginBody.Password {
		
		expirationTime := time.Now().Add(30 * time.Second)

		payload := vars.Payload {
			Email: LoginBody.Email,
			Password: LoginBody.Password,
			StandardClaims: jwt.StandardClaims{
				ExpiresAt: expirationTime.Unix(),
			},
		}

		token := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)

		tokenString, err := token.SignedString(jwtKey)

		if err != nil {
			log.Fatalf("Error in jwt token making: %v", err)
		}
		log.Println(tokenString)
		c.JSON(200, tokenString)

		con := databasepg.NewDBConn()
		defer con.Close()

		err = databasepg.InsertDB(con, LoginBody)

		if err != nil {
			log.Fatalf("error in inserting database: %v", err)
		}

	} else {
	c.JSON(404, "Password incorrect or you didn't signup yet")
	}

}