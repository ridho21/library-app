package model

import "time"

type Books struct {
	Id              string           `json:"id"`
	Title           string           `json:"title"`
	PublicationYear time.Time        `json:"publication_year"`
	Stock           int              `json:"stock"`
	TotalPages      int              `json:"total_pages"`
	CreatedAt       time.Time        `json:"created_at"`
	UpdatedAt       time.Time        `json:"updated_at"`
	Publisher       PublisherDetails `json:"publisher"`
	Author          AuthorDetails    `json:"author"`
	Category        Categories       `json:"category"`
}
