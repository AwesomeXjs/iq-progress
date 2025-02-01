package converter

import "github.com/AwesomeXjs/iq-progress/internal/model"

// ToTxData creates a new TxData instance.
func ToTxData(sender, receiver, amount int) *model.TxData {
	return &model.TxData{
		Sender:   sender,
		Receiver: receiver,
		Amount:   amount,
	}
}
