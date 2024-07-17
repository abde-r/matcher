package main

import (
	"log"
	"matchaVgo/cmd/api"
)

func main() {
	
	if err := api.Ga33ad_server(); err != nil {
		log.Fatal(err);
	}
}
