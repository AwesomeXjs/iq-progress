package service

import (
	"context"
	"testing"
	"time"

	"github.com/AwesomeXjs/iq-progress/internal/app"
	"github.com/AwesomeXjs/iq-progress/internal/model"
	"github.com/AwesomeXjs/iq-progress/internal/repository"
	service2 "github.com/AwesomeXjs/iq-progress/internal/service"
	"github.com/AwesomeXjs/iq-progress/pkg/logger"
	"github.com/AwesomeXjs/iq-progress/tests/unit/mocks"
	"github.com/gojuno/minimock/v3"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/require"
)

func TestGetOperation(t *testing.T) {
	t.Parallel()
	logger.Init(logger.GetCore(logger.GetAtomicLevel(app.LogLevel)))

	type IRepositoryMock func(mc *minimock.Controller) repository.IRepository

	type args struct {
		ctx    context.Context
		userID int
	}

	var (
		ctx = context.Background()
		mc  = minimock.NewController(t)

		userID  = 1
		timeNow = time.Now()

		response = []model.Operation{
			{
				ID:               1,
				FromUserID:       1,
				SenderUsername:   "user1",
				ToUserID:         2,
				ReceiverUsername: "user2",
				Amount:           100,
				Type:             "deposit",
				CreatedAt:        timeNow,
			},
		}
		someError = errors.New("some error")
	)
	defer t.Cleanup(mc.Finish)

	tests := []struct {
		name     string
		args     args
		want     []model.Operation
		err      error
		repoMock IRepositoryMock
	}{{
		name: "success",
		args: args{
			ctx:    ctx,
			userID: userID,
		},
		want: response,
		err:  nil,
		repoMock: func(mc *minimock.Controller) repository.IRepository {
			repo := mocks.NewIRepositoryMock(mc)
			repo.GetOperationsMock.Expect(ctx, userID).Return(response, nil)
			return repo
		},
	}, {
		name: "error",
		args: args{
			ctx:    ctx,
			userID: userID,
		},
		want: nil,
		err:  someError,
		repoMock: func(mc *minimock.Controller) repository.IRepository {
			repo := mocks.NewIRepositoryMock(mc)
			repo.GetOperationsMock.Expect(ctx, userID).Return(nil, someError)
			return repo
		},
	}}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			service := service2.New(tt.repoMock(mc), nil)

			res, err := service.GetOperations(tt.args.ctx, tt.args.userID)
			require.Equal(t, tt.err, err)
			require.Equal(t, tt.want, res)
		})
	}

}
