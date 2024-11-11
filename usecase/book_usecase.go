package usecase

import (
	"test-ordent/model"
	"test-ordent/repository"
	"time"

	"github.com/google/uuid"
)

type BookUsecase interface {
	GetAllBooks() ([]model.Books, error)
	InsertNewBook(payload model.Books) (model.Books, error)
	UpdateBooks(payload model.Books) error
	DeleteBooks(id string) error
}

type bookUsecase struct {
	repo repository.BookRepository
}

// DeleteBooks implements BookUsecase.
func (u *bookUsecase) DeleteBooks(id string) error {
	return u.repo.DeleteBook(id)
}

// UpdateBooks implements BookUsecase.
func (u *bookUsecase) UpdateBooks(payload model.Books) error {
	return u.repo.UpdateBook(payload)
}

// InsertNewBook implements BookUsecase.
func (u *bookUsecase) InsertNewBook(payload model.Books) (model.Books, error) {
	payload.Id = uuid.NewString()
	payload.CreatedAt = time.Now().UTC()
	return u.repo.InsertBook(payload)
}

func (u *bookUsecase) GetAllBooks() ([]model.Books, error) {
	return u.repo.GetAllBook()
}

func NewBookUsecase(
	repo repository.BookRepository,
) BookUsecase {
	return &bookUsecase{
		repo: repo,
	}
}
