package main

import (
	"casestudy/backend-repo/handler"
	"casestudy/backend-repo/repository"
	"casestudy/backend-repo/service"
	"context"
	"fmt"
	"log"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

func main() {
	err := NewApplication(8086)
	if err != nil {
		log.Fatal(err)
	}

}

func NewApplication(port int) error {

	log.SetFlags(log.LstdFlags | log.Lshortfile)

	mongoClient, err := mongo.NewClient(options.Client().ApplyURI("mongodb+srv://aykutisik:Ayk-0109@firstcluster.o4wm6.mongodb.net/todo_database?retryWrites=true&w=majority"))
	if err != nil {
		log.Fatal(err)
	}

	err = mongoClient.Connect(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	if err = mongoClient.Ping(ctx, readpref.Primary()); err != nil {
		fmt.Printf("could not ping to mongo db service: %v\n", err)
	}
	database := mongoClient.Database("todo_database")
	collection := database.Collection("todo_list_elements")
	repo := repository.NewRepository(database, mongoClient, collection)
	service := service.NewService(repo)
	handler := handler.NewHandler(service)

	app := NewServer(handler)

	return app.Listen(fmt.Sprintf(":%d", port))
}

func NewServer(handler_all handler.Handler) *fiber.App {

	app := fiber.New()
	app.Use(cors.New())
	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowHeaders: "Origin, Content-Type, Accept",
	}))

	app.Get("/GetTodoElements", handler_all.GetTodoElements)
	app.Post("/CreateTodo", handler_all.CreateTodo)
	app.Put("/DeleteAll", handler_all.DeleteAll)

	return app

}
