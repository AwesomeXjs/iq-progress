package service

import (
	"context"
	"testing"

	"github.com/AwesomeXjs/iq-progress/internal/app"
	"github.com/AwesomeXjs/iq-progress/internal/model"
	"github.com/AwesomeXjs/iq-progress/internal/service"
	"github.com/AwesomeXjs/iq-progress/pkg/dbclient"
	"github.com/AwesomeXjs/iq-progress/pkg/logger"
	"github.com/AwesomeXjs/iq-progress/tests/unit/mocks"
	"github.com/brianvoe/gofakeit"
	"github.com/gojuno/minimock/v3"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/require"
)

func TestSend(t *testing.T) {
	t.Parallel()
	logger.Init(logger.GetCore(logger.GetAtomicLevel(app.LogLevel)))

	type txManagerMock func(mc *minimock.Controller) dbclient.TxManager

	type args struct {
		ctx  context.Context
		data model.SendRequest
	}

	var (
		ctx context.Context
		mc  = minimock.NewController(t)

		senderID   = gofakeit.Int32()
		receiverID = gofakeit.Int32()
		amount     = gofakeit.Int8()

		req = model.SendRequest{
			Sender:   int(senderID),
			Receiver: int(receiverID),
			Amount:   int(amount),
		}

		someError = errors.New("some error")
	)
	defer t.Cleanup(mc.Finish)

	tests := []struct {
		name          string
		args          args
		want          int
		err           error
		txManagerMock txManagerMock
	}{
		{
			name: "success",
			args: args{
				ctx:  ctx,
				data: req,
			},
			want: 0,
			err:  nil,
			txManagerMock: func(mc *minimock.Controller) dbclient.TxManager {
				txManager := mocks.NewTxManagerMock(mc)
				txManager.ReadCommittedMock.Return(nil)

				return txManager
			},
		},
		{
			name: "error",
			args: args{
				ctx:  ctx,
				data: req,
			},
			want: 0,
			err:  someError,
			txManagerMock: func(mc *minimock.Controller) dbclient.TxManager {
				txManager := mocks.NewTxManagerMock(mc)
				txManager.ReadCommittedMock.Return(someError)

				return txManager
			},
		},
	}

	for _, tt := range tests {
		tt := tt

		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			svc := &service.Service{TxManager: tt.txManagerMock(mc)}

			res, err := svc.Send(tt.args.ctx, tt.args.data)
			require.Equal(t, tt.err, err)
			require.Equal(t, tt.want, res)
		})
	}
}
