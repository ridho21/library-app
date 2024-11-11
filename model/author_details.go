package model

type AuthorDetails struct {
	Id          string `json:"id"`
	AuthorName  string `json:"author_name"`
	Email       string `json:"email"`
	PhoneNumber string `json:"phone_number"`
}
