package response

import "time"

type LoginResponseDto struct {
	AccessToken string `json:"accessToken"`
	UserId      string `json:"userId"`
}

type RegisterResponseDto struct {
	Id          string    `json:"id"`
	FullName    string    `json:"full_name"`
	Email       string    `json:"email"`
	PhoneNumber string    `json:"phone_number"`
	Address     string    `json:"address"`
	Role        string    `json:"role"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
