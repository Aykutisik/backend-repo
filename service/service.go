package service

import (
	"casestudy/backend-repo/model"
	"casestudy/backend-repo/repository"
)

type Service interface {
	CreateTodo(todo model.SendTodoElements) (model.SendTodoElements, error)
	GetTodoElements() (todos []model.TodoElements, err error)
	DeleteAll(theList []model.TodoElements) ([]model.TodoElements, error)
}

type service struct {
	repo repository.Repository
}

var _ Service = service{}

func NewService(repo repository.Repository) Service {
	return service{repo: repo}
}

func (s service) GetTodoElements() (todos []model.TodoElements, err error) {
	return s.repo.GetTodoElements()
}

func (s service) CreateTodo(todo model.SendTodoElements) (model.SendTodoElements, error) {
	return s.repo.CreateTodo(todo)
}

func (s service) DeleteAll(theList []model.TodoElements) ([]model.TodoElements, error) {
	return s.repo.DeleteAll(theList)
}
