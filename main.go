package main

import (
	"fmt"
	"log"
	"os"
	"zaidanweb/config"
	"zaidanweb/httpserver/controllers"
	"zaidanweb/httpserver/repositories"
	"zaidanweb/httpserver/routers"
	"zaidanweb/httpserver/services"
	"zaidanweb/utils"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

/// @contact.name  API Support
// @contact.url   http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url  http://www.apache.org/licenses/LICENSE-2.0.html

// @host                       localhost:3030
// @BasePath                   /api
// @securityDefinitions.apikey BearerAuth
// @in                         header
// @name                       Authorization
func main() {
	err := godotenv.Load(".env")
	fmt.Println("PGHOST", os.Getenv("PGHOST"))

	if err != nil {
		log.Fatal("Environment Variables not found")
	}
	app := gin.Default()
	appRoute := app.Group("/api")
	db, _ := config.Connect()

	authService := utils.NewAuthHelper(utils.Constants.JWT_SECRET_KEY)

	userRepository := repositories.NewUserRepository(db)
	userService := services.NewUserService(userRepository)
	userController := controllers.NewUserController(userService, authService)

	routers.UserRouter(appRoute, userController, authService)

	app.Run(":3030")
}
