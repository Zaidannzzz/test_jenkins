package main

import (
	"backend/config"
	"backend/httpserver/controllers"
	"backend/httpserver/repositories"
	"backend/httpserver/routers"
	"backend/httpserver/services"
	"backend/utils"
	"log"

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
	if err != nil {
		log.Fatal("Environment Variables not found")
	}

	gin.SetMode(gin.ReleaseMode)

	app := gin.Default()
	appRoute := app.Group("/api")

	db, err := config.Connect()
	if err != nil {
		log.Fatal("Failed to connect to the database")
	}

	authService := utils.NewAuthHelper(utils.Constants.JWT_SECRET_KEY)

	userRepository := repositories.NewUserRepository(db)
	userService := services.NewUserService(userRepository)
	userController := controllers.NewUserController(userService, authService)

	routers.UserRouter(appRoute, userController, authService)

	err = app.Run()
	if err != nil {
		log.Fatal("Failed to start the server")
	}
}
