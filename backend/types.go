package main

import "time"

type User struct {
	ID        	int64  	`json:"id"`
	FirstName	string  	`json:"firstName"`
	LastName	string  	`json:"lastName"`
	Password	string  	`json:"password"`
	CreatedAt time.Time `json:"createdAt"`
}

type Post struct {
	ID        int64  	`json:"id"`
	Title      string 	`json:"title"`
	CreatedAt time.Time `json:"createdAt"`
}
