package main

type User struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

type Message struct {
	Id      string
	Message string
	UserId  string
}
