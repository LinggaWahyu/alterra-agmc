package factory

import (
	"alterra-agmc-day-5-6/database"
	"alterra-agmc-day-5-6/internal/repository"
)

type Factory struct {
	BookRepository repository.Book
	UserRepository repository.User
}

func NewFactory() *Factory {
	db := database.GetConnection()
	return &Factory{
		BookRepository: repository.NewBook(db),
		UserRepository: repository.NewUser(db),
	}
}
