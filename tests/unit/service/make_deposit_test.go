package service

import (
	"context"
	"testing"

	"github.com/AwesomeXjs/iq-progress/internal/app"
	"github.com/AwesomeXjs/iq-progress/internal/model"
	service2 "github.com/AwesomeXjs/iq-progress/internal/service"
	"github.com/AwesomeXjs/iq-progress/pkg/dbclient"
	"github.com/AwesomeXjs/iq-progress/pkg/logger"
	"github.com/AwesomeXjs/iq-progress/tests/unit/mocks"
	"github.com/brianvoe/gofakeit"
	"github.com/gojuno/minimock/v3"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/require"
)

func MakeDepositTest(t *testing.T) {
	t.Parallel()
	logger.Init(logger.GetCore(logger.GetAtomicLevel(app.LogLevel)))

	type txManagerMock func(mc *minimock.Controller) dbclient.TxManager

	type args struct {
		ctx  context.Context
		data model.DepositRequest
	}

	var (
		ctx = context.Background()
		mc  = minimock.NewController(t)

		userID = gofakeit.Int32()
		amount = gofakeit.Int8()

		req = model.DepositRequest{
			UserID: int(userID),
			Amount: int(amount),
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
				txMMock := mocks.NewTxManagerMock(mc)
				txMMock.ReadCommittedMock.Return(nil)
				return txMMock
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
				txMMock := mocks.NewTxManagerMock(mc)
				txMMock.ReadCommittedMock.Return(someError)
				return txMMock
			},
		},
	}

	for _, tt := range tests {
		tt := tt

		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			service := &service2.Service{TxManager: tt.txManagerMock(mc)}
			res, err := service.MakeDeposit(tt.args.ctx, tt.args.data)

			require.Equal(t, tt.err, err)
			require.Equal(t, tt.want, res)
		})
	}
}
