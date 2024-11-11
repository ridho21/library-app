package model

import "time"

type TransactionBooks struct {
	Id                 string                    `json:"id"`
	Status             string                    `json:"status"`
	Penalty            int                       `json:"penalty"`
	BorrowDate         time.Time                 `json:"borrow_date"`
	ReturnDate         time.Time                 `json:"return_date"`
	ReturnActual       time.Time                 `json:"return_actual"`
	CreatedAt          time.Time                 `json:"created_at"`
	UpdatedAt          time.Time                 `json:"updated_at"`
	UserId             string                    `json:"user_id"`
	TransactionDetails []TransactionBooksDetails `json:"transaction_details"`
}

type TransactionBooksDetails struct {
	Id            string `json:"id"`
	TransactionId string `json:"transaction_id"`
	BooksId       string `json:"books_id"`
	Qty           int    `json:"qty"`
}
