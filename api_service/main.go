package main

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	"gitlab.com/gurugin/configs"
	"gitlab.com/gurugin/db"
	"gitlab.com/gurugin/handlers"
	"gitlab.com/gurugin/repositories"
	"gitlab.com/gurugin/routers"
	"gitlab.com/gurugin/services"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
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
// @BasePath /
func main() {
	configs.InitialEnv("../.env")
	db := db.NewPostgresDatabase()
	app := fiber.New()

	// Use grpc.WithTransportCredentials with insecure.NewCredentials() for an insecure connection
	conn, err := grpc.NewClient(fmt.Sprintf("%s:%s", configs.GetMLHOST(), configs.GetMLPort()), grpc.WithTransportCredentials(insecure.NewCredentials()), grpc.WithDefaultCallOptions(grpc.MaxCallSendMsgSize(1024*1024*50), grpc.MaxCallRecvMsgSize(1024*1024*50)))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	recipeRepository := repositories.NewMLRepository(db)
	mlClient := services.NewMLServiceClient(conn)
	mlService := services.NewMLService(mlClient, recipeRepository)
	routers.SetupRoutes(app.Group("/api"), handlers.NewMLHandler(mlService))

	if err := app.Listen(fmt.Sprintf(":%s", configs.GetAPIPort())); err != nil {
		log.Fatal(err)
	}

}
