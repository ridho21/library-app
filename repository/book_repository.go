package repository

import (
	"database/sql"
	"errors"
	"test-ordent/model"
)

type BookRepository interface {
	GetAllBook() ([]model.Books, error)
	// GetBookByid(id string) (model.Books, error)
	InsertBook(payload model.Books) (model.Books, error)
	UpdateBook(payload model.Books) error
	DeleteBook(id string) error
}

type bookRepository struct {
	db *sql.DB
}

func (u *bookRepository) GetAllBook() ([]model.Books, error) {
	var books []model.Books
	rows, err := u.db.Query(`SELECT id, title, publication_year, stock, total_pages, created_at, updated_at, publisher_id, author_id, category_id from mst_books`)

	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var book model.Books
		var publisherDetails model.PublisherDetails
		var authorDetails model.AuthorDetails
		var categoryDetails model.Categories

		err := rows.Scan(
			&book.Id,
			&book.Title,
			&book.PublicationYear,
			&book.Stock,
			&book.TotalPages,
			&book.CreatedAt,
			&book.UpdatedAt,
			&book.Publisher.Id,
			&book.Author.Id,
			&book.Category.Id,
		)

		if err != nil {
			return nil, err
		}

		rowPublisherDetails, err := u.db.Query(`SELECT * FROM MST_PUBLISHER_DETAILS WHERE id = $1`, book.Publisher.Id)

		if err != nil {
			return nil, errors.New("publisher")
		}

		for rowPublisherDetails.Next() {
			var publisherDetail model.PublisherDetails

			err = rowPublisherDetails.Scan(
				&publisherDetail.Id,
				&publisherDetail.PublisherName,
				&publisherDetail.Email,
				&publisherDetail.PhoneNumber,
			)

			if err != nil {
				return nil, errors.New("publisher")
			}

			publisherDetails = publisherDetail
		}

		rowAuthorDetails, err := u.db.Query(`SELECT * FROM MST_AUTHOR_DETAILS WHERE id = $1`, book.Author.Id)

		if err != nil {
			return nil, errors.New("author")
		}

		for rowAuthorDetails.Next() {
			var authorDetail model.AuthorDetails

			err = rowAuthorDetails.Scan(
				&authorDetail.Id,
				&authorDetail.AuthorName,
				&authorDetail.Email,
				&authorDetail.PhoneNumber,
			)

			if err != nil {
				return nil, errors.New("author")
			}

			authorDetails = authorDetail
		}

		rowCategories, err := u.db.Query(`SELECT * FROM MST_CATEGORIES WHERE id = $1`, book.Category.Id)

		if err != nil {
			return nil, errors.New("categories")
		}

		for rowCategories.Next() {
			var categoryDetail model.Categories

			err = rowCategories.Scan(
				&categoryDetail.Id,
				&categoryDetail.CategoryName,
			)

			if err != nil {
				return nil, errors.New("categories")
			}

			categoryDetails = categoryDetail
		}

		book.Publisher = publisherDetails
		book.Author = authorDetails
		book.Category = categoryDetails
		books = append(books, book)
	}
	return books, nil
}

// func (u *bookRepository) GetBookByid(id string) (model.Books, error) {
// 	var books model.Books
// 	err := u.db.QueryRow(`SELECT id, title, publication_year, stock, total_pages, created_at, updated_at, publisher_id, author_id, category_id from mst_books WHERE id=$1`, id).Scan(
// 		&books.Id,
// 	)
// 	if err != nil {
// 		return model.Books{}, err
// 	}
// 	return books, nil
// }

func (u *bookRepository) InsertBook(payload model.Books) (model.Books, error) {
	var book model.Books
	err := u.db.QueryRow(`INSERT INTO mst_books (id, title, publication_year, stock, total_pages, created_at, updated_at, publisher_id, author_id, category_id) VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9,$10) RETURNING id, title, publication_year, stock, total_pages, created_at, updated_at, publisher_id, author_id, category_id`,
		payload.Id,
		payload.Title,
		payload.PublicationYear,
		payload.Stock,
		payload.TotalPages,
		payload.CreatedAt,
		payload.UpdatedAt,
		payload.Publisher.Id,
		payload.Author.Id,
		payload.Category.Id,
	).Scan(
		&book.Id,
		&book.Title,
		&book.PublicationYear,
		&book.Stock,
		&book.TotalPages,
		&book.CreatedAt,
		&book.UpdatedAt,
		&book.Publisher.Id,
		&book.Author.Id,
		&book.Category.Id,
	)
	if err != nil {
		return model.Books{}, err
	}
	return book, nil
}

func (u *bookRepository) UpdateBook(payload model.Books) error {
	_, err := u.db.Exec(`UPDATE MST_BOOKS SET title = $1, publication_year = $2, stock = $3, total_pages = $4 WHERE id = $5`,
		payload.Title,
		payload.PublicationYear,
		payload.Stock,
		payload.TotalPages,
		payload.Id,
	)
	if err != nil {
		return err
	}
	return nil
}

func (u *bookRepository) DeleteBook(id string) error {
	_, err := u.db.Exec(`DELETE FROM MST_BOOKS WHERE id = $1`, id)
	if err != nil {
		return err
	}
	return nil
}

func NewBookRepository(db *sql.DB) BookRepository {
	return &bookRepository{
		db: db,
	}
}
