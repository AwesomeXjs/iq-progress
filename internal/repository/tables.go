package repository

const (
	UserTable             = "users"
	IDColumn              = "id"
	TransactionTable      = "transactions"
	AmountColumn          = "amount"
	FromUserIDColumn      = "from_user_id"
	ToUserIDColumn        = "to_user_id"
	TransactionDataColumn = "transaction_data"
	BalanceColumn         = "balance"
	TypeColumn            = "type"

	ReturnBalanceColumn = "RETURNING balance"

	// Aliases
	TxTable          = "transactions t"
	TxID             = "t.id AS transaction_id"
	TxType           = "t.type AS transaction_type"
	TxAmount         = "t.amount"
	TxDate           = "t.transaction_date"
	U1               = "users u1"
	U2               = "users u2"
	SenderID         = "t.from_user_id AS sender_id"
	SenderUsername   = "u1.username AS sender_username"
	ReceiverID       = "t.to_user_id AS receiver_id"
	ReceiverUsername = "u2.username AS receiver_username"
)
