package main

import "time"

type User struct {
	Id        	int64  		`json:"id"`
	FirstName	string  	`json:"firstName"`
	LastName	string  	`json:"lastName"`
	Username	string  	`json:"username"`
	Email		string  	`json:"email"`
	Password	string  	`json:"password"`
	Gender		string  	`json:"gender"`
	CreatedAt 	time.Time	`json:"createdAt"`
}

type Post struct {
	Id        int64  	`json:"id"`
	Title      string 	`json:"title"`
	CreatedAt time.Time `json:"createdAt"`
}
