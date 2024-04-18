package model

type UserCredentials struct{
	Email 	 string `json:"email"`
	Password string `json:"password"`
}

type UserSession struct{
	JWTToken string	`json:"jwt_token"`
}