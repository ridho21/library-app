package model

type PublisherDetails struct {
	Id            string `json:"id"`
	PublisherName string `json:"publisher_name"`
	Email         string `json:"email"`
	PhoneNumber   string `json:"phone_number"`
}
