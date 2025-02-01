package model

import "time"

type Operations struct {
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
	Sender   int `json:"sender" valid:"required"`
	Receiver int `json:"receiver" valid:"required"`
	Amount   int `json:"amount" valid:"required,range(1|100000000)"`
}

type DepositRequest struct {
	UserID int `json:"user_id" valid:"required"`
	Amount int `json:"amount" valid:"required,range(1|100000000)"`
}
