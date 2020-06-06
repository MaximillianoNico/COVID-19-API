package main

import (
	"flag"
	"github.com/MaximillianoNico/COVID-19-API/internals/httpserver/routers"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"os"
)

func LoadEnv() {
	mode := flag.String("mode", "development", "mode go run server")
	flag.Parse()

	if *mode == "production" || *mode == "staging" {
		err := godotenv.Load(".env." + *mode)

		if err != nil {
			log.Fatal("Error loading .env file")
		}
	} else {
		err := godotenv.Load()

		if err != nil {
			log.Fatal("Error loading .env file")
		}
	}
}

// @title Swagger Example API
// @version 1.0
// @description This is a sample server celler server.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8081
// @BasePath /api
// @query.collection.format multi

// @securityDefinitions.basic BasicAuth

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization

// @securitydefinitions.oauth2.application OAuth2Application
// @tokenUrl https://example.com/oauth/token
// @scope.write Grants write access
// @scope.admin Grants read and write access to administrative information

// @securitydefinitions.oauth2.implicit OAuth2Implicit
// @authorizationurl https://example.com/oauth/authorize
// @scope.write Grants write access
// @scope.admin Grants read and write access to administrative information

// @securitydefinitions.oauth2.password OAuth2Password
// @tokenUrl https://example.com/oauth/token
// @scope.read Grants read access
// @scope.write Grants write access
// @scope.admin Grants read and write access to administrative information

// @securitydefinitions.oauth2.accessCode OAuth2AccessCode
// @tokenUrl https://example.com/oauth/token
// @authorizationurl https://example.com/oauth/authorize
// @scope.admin Grants read and write access to administrative information

// @x-extension-openapi {"example": "value on a json format"}

func main() {
	LoadEnv() // load environment

	routersInit := routers.InitRouters() // init all endpoint
	port := ":" + os.Getenv("API_PORT")

	server := &http.Server{
		Addr:    port,
		Handler: routersInit,
	}

	log.Printf("[info] start http server listening %s", os.Getenv("API_PORT"))

	err := server.ListenAndServe()
	if err != nil {
		log.Printf("Server err: %v", err)
	}

}
