package model

// TxData represents transaction data including sender, receiver, and amount.
type TxData struct {
	Sender   int `json:"sender" valid:"required"`
	Receiver int `json:"receiver" valid:"required"`
	Amount   int `json:"amount" valid:"required,range(1|100000000)"`
}
