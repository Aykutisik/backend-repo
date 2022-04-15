package repository

import (
	"casestudy/backend-repo/model"
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type Repository interface {
	CreateTodo(todo model.SendTodoElements) (model.SendTodoElements, error)
	GetTodoElements() (todos []model.TodoElements, err error)
	DeleteAll(theList []model.TodoElements) ([]model.TodoElements, error)
}

type repository struct {
	db                     *mongo.Database
	mongoClient            *mongo.Client
	TodoElementsCollection *mongo.Collection
}

var _ Repository = repository{}

func NewRepository(db *mongo.Database, mongoClient *mongo.Client, TodoElementsCollection *mongo.Collection) Repository {
	return repository{db: db, mongoClient: mongoClient, TodoElementsCollection: TodoElementsCollection}
}

func (r repository) GetTodoElements() (todos []model.TodoElements, err error) {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	cursor, err := r.TodoElementsCollection.Find(ctx, bson.M{})
	if err != nil {
		log.Fatal(err)
	}

	var elements []bson.M
	if err = cursor.All(ctx, &todos); err != nil {
		log.Fatal(err)
	}

	bsonBytes, _ := bson.Marshal(elements)
	bson.Unmarshal(bsonBytes, &todos)

	return todos, err
}

func (r repository) CreateTodo(todo model.SendTodoElements) (model.SendTodoElements, error) {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

	r.TodoElementsCollection.InsertOne(ctx, todo)

	return todo, nil
}

func (r repository) DeleteAll(theList []model.TodoElements) ([]model.TodoElements, error) {

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	arraylen := len(theList)
	for i := 0; i < arraylen; i++ {
		id := theList[i].Id
		fmt.Println(id)
		//idPrimitive, _ := primitive.ObjectIDFromHex(id)

		_, err := r.TodoElementsCollection.DeleteOne(ctx, bson.M{"_id": id})
		if err != nil {
			return theList, err
		}
	}

	return theList, nil
}
