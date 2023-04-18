package uow

import (
	"example/db"
	"example/repository"
)

type UnitOfWork interface {
	Todos() repository.TodoRepository
}

type unitOfWork struct {
	db *db.Database
}

func NewUnitOfWork(db *db.Database) UnitOfWork {
	return &unitOfWork{db: db}
}

func (uow *unitOfWork) Todos() repository.TodoRepository {
	return repository.NewTodoRepository(uow.db)
}
