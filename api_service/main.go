package main

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	"gitlab.com/gurugin/handlers"
	"gitlab.com/gurugin/models"
	"gitlab.com/gurugin/routers"
	"gitlab.com/gurugin/services"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	// replace with your own docs folder, usually "github.com/username/reponame/docs"
	// _ "github.com/gofiber/swagger/example/docs"
)

// @title Fiber Example API
// @version 1.0
// @description This is a sample swagger for Fiber
// @termsOfService http://swagger.io/terms/
// @contact.name API Support
// @contact.email fiber@swagger.io
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @host localhost:8080
// @BasePath /
func main() {
	dsn := "root:secret@tcp(127.0.0.1:3306)/gurugindb?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	err = db.AutoMigrate(
		&models.User{},
		&models.Menu{},
		&models.Ingredient{},
		&models.IngredientOnMenu{},
		&models.Recipe{},
		&models.Comment{},
		&models.Favorite{},
		&models.Notification{},
	)
	if err != nil {
		log.Fatal(err)
	}

	app := fiber.New()

	// Use grpc.WithTransportCredentials with insecure.NewCredentials() for an insecure connection
	conn, err := grpc.NewClient("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	mlClient := services.NewMLServiceClient(conn)
	routers.SetupRoutes(app.Group("/api"), handlers.NewMLHandler(mlClient))

	if err := app.Listen(fmt.Sprintf(":%s", "8080")); err != nil {
		log.Fatal(err)
	}

}
