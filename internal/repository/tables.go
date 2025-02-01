package repository

// Constants for database table names, columns, and aliases.
const (
	// UserTable is the name of the users table.
	UserTable = "users"
	// IDColumn is the column name for the user ID.
	IDColumn = "id"
	// TransactionTable is the name of the transactions table.
	TransactionTable = "transactions"
	// AmountColumn is the column name for the transaction amount.
	AmountColumn = "amount"
	// FromUserIDColumn is the column name for the sender's user ID.
	FromUserIDColumn = "from_user_id"
	// ToUserIDColumn is the column name for the receiver's user ID.
	ToUserIDColumn = "to_user_id"
	// BalanceColumn is the column name for the user's balance.
	BalanceColumn = "balance"
	// TypeColumn is the column name for the transaction type.
	TypeColumn = "type"

	// ReturnBalanceColumn is used to return the updated balance after a transaction.
	ReturnBalanceColumn = "RETURNING balance"

	// Aliases

	// TxTable is an alias for the transactions table.
	TxTable = "transactions t"
	// TxID is an alias for the transaction ID.
	TxID = "t.id AS transaction_id"
	// TxType is an alias for the transaction type.
	TxType = "t.type AS transaction_type"
	// TxAmount is an alias for the transaction amount.
	TxAmount = "t.amount"
	// TxDate is an alias for the transaction date.
	TxDate = "t.transaction_date"
	// U1 is an alias for the users table (sender).
	U1 = "users u1"
	// U2 is an alias for the users table (receiver).
	U2 = "users u2"
	// SenderID is an alias for the sender's user ID.
	SenderID = "t.from_user_id AS sender_id"
	// SenderUsername is an alias for the sender's username.
	SenderUsername = "u1.username AS sender_username"
	// ReceiverID is an alias for the receiver's user ID.
	ReceiverID = "t.to_user_id AS receiver_id"
	// ReceiverUsername is an alias for the receiver's username.
	ReceiverUsername = "u2.username AS receiver_username"
)
