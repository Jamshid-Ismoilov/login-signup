package vars

import (
	jwt "github.com/dgrijalva/jwt-go"
)
type LoginBody struct {
	Email, Password string
}

type SignupBody struct {
	Email string
}

type ChangePasswordBody struct {
	Email, CurrentPassword, NewPassword string
}


type Payload struct {
	Email, Password string
	jwt.StandardClaims
}