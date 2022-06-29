package models

type User struct {
	UserID       string `json:"userid"`
	UserName     string `json:"username"`
	UserPassword string `json:"userpassword"`
}
