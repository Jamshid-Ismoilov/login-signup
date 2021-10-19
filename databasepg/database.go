package databasepg

import (
	"fmt"
	"log"

	"app/vars"

	"github.com/go-pg/pg"
	"database/sql"
	_"github.com/lib/pq"
	"app/config"
)

var SELECT = `select email from login_bodies where email = $1`

var SELECTPE = `select email, password from login_bodies where email = $1`

var UPDATE = `UPDATE login_bodies SET password = $2 WHERE email = $1`

func NewDBConn() (con *pg.DB) {
   address := fmt.Sprintf("%s:%s", "localhost", "5432")
   options := &pg.Options{
      User:     "jamshid",
      Password: "1111",
      Addr:     address,
      Database: "login",
      PoolSize: 50,
   }

   con = pg.Connect(options)
   if con == nil {
      log.Fatal("cannot connect to postgres")
   }
	return
}

func InsertDB(pg *pg.DB, post vars.LoginBody) error {
   _, err := pg.Model(&post).Insert()
   return err
}

func SelectDBPost(post vars.SignupBody) (string, error) {
   conn, err := sql.Open("postgres", config.DB_CONFIG)
   defer conn.Close()
   if err != nil {
   	log.Fatalf("error in connecting psql: %v", err)
   }
   var result string
   err = conn.QueryRow(SELECT, &post.Email,).Scan(&result)

   return result, err
}


func ChangePasswordDB(post vars.ChangePasswordBody) int {
	
	conn, err := sql.Open("postgres", config.DB_CONFIG)
	defer conn.Close()

	if err != nil {
		log.Fatalf("error in connecting psql change password: %v", err)
	}

	var checkpost vars.LoginBody
	
	err = conn.QueryRow(SELECTPE, &post.Email).Scan(&checkpost.Email, &checkpost.Password)
	if err != nil {
		log.Fatalf("error in selectpe query: %v", err)
	}

	if checkpost.Email == post.Email && checkpost.Password == post.CurrentPassword {
		_ = conn.QueryRow(UPDATE, &post.Email, &post.NewPassword)
		return 1
	}
	return 0
}