package model

import "time"

type Operation struct {
	ID               int       `json:"id" db:"transaction_id"`
	FromUserID       int       `json:"from_user_id" db:"sender_id"`
	SenderUsername   string    `json:"sender_username" db:"sender_username"`
	ToUserID         int       `json:"to_user_id" db:"receiver_id"`
	ReceiverUsername string    `json:"receiver_username" db:"receiver_username"`
	Amount           int       `json:"amount" db:"amount"`
	Type             string    `json:"type" db:"transaction_type"`
	CreatedAt        time.Time `json:"transaction_date" db:"transaction_date"`
}

type SendRequest struct {
	Sender   int `json:"sender" valid:"required" example:"1"`
	Receiver int `json:"receiver" valid:"required" example:"2"`
	Amount   int `json:"amount" valid:"required,range(1|100000000)" example:"100"`
}

type DepositRequest struct {
	UserID int `json:"user_id" valid:"required" example:"1"`
	Amount int `json:"amount" valid:"required,range(1|100000000)" example:"1000"`
}
