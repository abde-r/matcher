package main

import (
	"log"
	"matcher/cmd/api"
	_ "matcher/docs"
)

// @title           matcher API
// @version         1.0
// @description     This is a sample server celler server.
// @termsOfService  http://swagger.io/terms/

// @contact.name   API Support
// @contact.url    http://www.swagger.io/support
// @contact.email  support@swagger.io

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      localhost:8000
// @BasePath  /api/v1

// @securityDefinitions.basic  BasicAuth

// @externalDocs.description  matcherAPI
// @externalDocs.url          https://swagger.io/resources/open-api/
func main() {
	
	if err := api.Ga33ad_server(); err != nil {
		log.Fatal(err);
	}
}
