package config

import (
  "fmt"
  "github.com/go-redis/redis"
)


var (
  host = "localhost"
  user = "jamshid"
  password = "1111"
  dbname = "login"
  port = 5432
)

var DB_CONFIG = fmt.Sprintf(
  "host=%s user=%s password=%s dbname=%s port=%d", 
  host, user, password, dbname, port,
)

 var GSender = "gamerdeveloper1408@gmail.com"
 var GPassword = "gamerdev.1.1.1"
 var GHost = "smtp.gmail.com"
 var GPort = "587"

var Client = redis.NewClient(&redis.Options{
	Addr: "localhost:6379",
	Password: "jubajuba",
	DB: 0,	
})