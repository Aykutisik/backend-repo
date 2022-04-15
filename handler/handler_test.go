package handler

import (
	"bytes"
	"casestudy/backend-repo/model"
	"casestudy/backend-repo/service"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/golang/mock/gomock"
	"go.mongodb.org/mongo-driver/bson/primitive"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func TestGetTodoElements(t *testing.T) {

	t.Run("Database connection should be established.",
		func(t *testing.T) {
			mongoClient, err := mongo.NewClient(options.Client().ApplyURI("mongodb+srv://aykutisik:Ayk-0109@firstcluster.o4wm6.mongodb.net/todo_database?retryWrites=true&w=majority"))
			assert.Nil(t, err)

			err = mongoClient.Connect(context.Background())
			assert.Nil(t, err)
		})

	t.Run("GetTodoElements should response with status code 200",
		func(t *testing.T) {
			serviceMockController := gomock.NewController(t)
			service := service.NewMockService(serviceMockController)
			handler := NewHandler(service)
			service.EXPECT().GetTodoElements().Return(getAllTodoElements(), nil)
			app := fiber.New()
			app.Get("/GetTodoElements", handler.GetTodoElements)
			req := httptest.NewRequest("GET", fmt.Sprintf("/GetTodoElements"), nil)
			res, err := app.Test(req)
			assert.Nil(t, err)
			assert.Equal(t, 200, res.StatusCode)

		})

	t.Run("When GetTodoElements failed",
		func(t *testing.T) {
			serviceMockController := gomock.NewController(t)
			service := service.NewMockService(serviceMockController)
			handler := NewHandler(service)
			service.EXPECT().GetTodoElements().Return(getAllTodoElements(), errors.New("An error occured"))
			app := fiber.New()
			app.Get("/GetTodoElements", handler.GetTodoElements)
			req := httptest.NewRequest("GET", fmt.Sprintf("/GetTodoElements"), nil)
			res, err := app.Test(req)
			assert.Nil(t, err)
			assert.Equal(t, 400, res.StatusCode)
		})
}

func TestCreateTodo(t *testing.T) {

	t.Run("Database connection should be established.",
		func(t *testing.T) {
			mongoClient, err := mongo.NewClient(options.Client().ApplyURI("mongodb://mongoadmin:secret@localhost:27017/?authSource=admin"))
			assert.Nil(t, err)

			err = mongoClient.Connect(context.Background())
			assert.Nil(t, err)
		})

	t.Run("CreateTodo should response with status code 201",
		func(t *testing.T) {
			serviceMockController := gomock.NewController(t)
			service := service.NewMockService(serviceMockController)
			handler := NewHandler(service)
			testBody := model.SendTodoElements{Text: "test", Status: 0}
			testb := fiber.Map{"text": "test", "status": 0}
			requestByte, _ := json.Marshal(testb)
			requestReader := bytes.NewReader(requestByte)
			service.EXPECT().CreateTodo(testBody).Return(testBody, nil)

			app := fiber.New()
			app.Post("/CreateTodo", handler.CreateTodo)

			req := httptest.NewRequest(http.MethodPost, fmt.Sprintf("/CreateTodo"), requestReader)
			req.Header.Set("Content-Type", "application/json")
			res, _ := app.Test(req)
			assert.Equal(t, 201, res.StatusCode)
		})

	// t.Run("Cannot request with empty text",
	// 	func(t *testing.T) {
	// 		mongoClient, err := mongo.NewClient(options.Client().ApplyURI("mongodb://mongoadmin:secret@localhost:27017/?authSource=admin"))
	// 		assert.Nil(t, err)

	// 		err = mongoClient.Connect(context.Background())
	// 		assert.Nil(t, err)

	// 		database := mongoClient.Database("todo_database")
	// 		collection := database.Collection("todo_list_elements")

	// 		repo := repository.NewRepository(database, mongoClient, collection)
	// 		service := service.NewService(repo)
	// 		handler := NewHandler(service)

	// 		app := fiber.New()

	// 		app.Post("/CreateTodo", handler.CreateTodo)
	// 		testBody := model.SendTodoElements{Text: "", Status: 0}

	// 		requestByte, _ := json.Marshal(testBody)
	// 		requestReader := bytes.NewReader(requestByte)

	// 		req := httptest.NewRequest(http.MethodPost, fmt.Sprintf("/CreateTodo"), requestReader)
	// 		req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	// 		res, err := app.Test(req)
	// 		assert.Nil(t, err)
	// 		assert.Equal(t, 201, res.StatusCode)

	// 	})

}

func TestDeleteAll(t *testing.T) {

	t.Run("DeleteAll Test When ",
		func(t *testing.T) {
			serviceMockController := gomock.NewController(t)
			service := service.NewMockService(serviceMockController)
			handler := NewHandler(service)
			testBody := getAllTodoElements()
			//testb := fiber.Map{}
			const addressRequest = `[{"_id":"62234346c2a65768f2c03ca5","text":"drink water","status":0},{"_id":"6225bfaec2a65768f2c03ca6","text":"bla bla","status":0}]`
			//	requestByte, _ := json.Marshal(testb)
			//	requestReader := bytes.NewReader(requestByte)
			service.EXPECT().DeleteAll(testBody).Return(testBody, nil)

			app := fiber.New()
			app.Put("/DeleteAll", handler.DeleteAll)

			req := httptest.NewRequest(http.MethodPut, fmt.Sprintf("/DeleteAll"), strings.NewReader(addressRequest))
			req.Header.Set("Content-Type", "application/json")
			res, _ := app.Test(req)
			assert.Equal(t, 200, res.StatusCode)
		})

}

func getAllTodoElements() []model.TodoElements {
	var list []model.TodoElements
	//	[{ "_id": "62234346c2a65768f2c03ca5", "status": 0, "text": "drink water" }, { "_id": "6225bfaec2a65768f2c03ca6", "status": 0, "text": "bla bla" }]

	var item model.TodoElements

	item.Id, _ = primitive.ObjectIDFromHex("62234346c2a65768f2c03ca5")
	item.Text = "drink water"
	item.Status = 0

	list = append(list, item)

	item.Id, _ = primitive.ObjectIDFromHex("6225bfaec2a65768f2c03ca6")
	item.Text = "bla bla"
	item.Status = 0

	list = append(list, item)

	return list
}
