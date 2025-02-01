package converter

import "github.com/AwesomeXjs/iq-progress/internal/model"

func ToTxData(sender, receiver, amount int) *model.TxData {
	return &model.TxData{
		Sender:   sender,
		Receiver: receiver,
		Amount:   amount,
	}
}
