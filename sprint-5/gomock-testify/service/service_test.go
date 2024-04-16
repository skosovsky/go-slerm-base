package service_test

import (
	"github.com/skosovsky/go-slerm-base/sprint-5/gomock-testify/service"
	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"
	"testing"
)

func TestService(t *testing.T) {
	testCases := map[string]struct {
		username           string
		amount             int
		balance            int
		initPaymentGateway func(m *service.MockPaymentGateway)
		initLogger         func(m *service.MockLogger)
		err                error
	}{
		"success": {
			username: "Jack",
			amount:   1,
			balance:  10,
			initPaymentGateway: func(m *service.MockPaymentGateway) {
				m.EXPECT().SendMoneyAndGetCurrentBalance("Jack", 1).Return(10, nil)
			},
			initLogger: func(m *service.MockLogger) {
				m.EXPECT().Info("send money ok")
			},
			err: nil,
		},
		"error": {
			username: "John",
			amount:   10,
			balance:  0,
			initPaymentGateway: func(m *service.MockPaymentGateway) {
				m.EXPECT().SendMoneyAndGetCurrentBalance("John", 10).Return(0, service.ErrConnection)
			},
			initLogger: func(m *service.MockLogger) {
				m.EXPECT().Error("send money error")
			},
			err: service.ErrConnection,
		},
	}

	for name, testCase := range testCases {
		t.Run(name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			extDep := service.NewMockPaymentGateway(ctrl)
			if testCase.initPaymentGateway != nil {
				testCase.initPaymentGateway(extDep)
			}
			logger := service.NewMockLogger(ctrl)
			if testCase.initLogger != nil {
				testCase.initLogger(logger)
			}
			balance, err := service.TransferMoney(extDep, logger, testCase.username, testCase.amount)
			if err != nil {
				require.ErrorIs(t, err, testCase.err)
			} else {
				require.NoError(t, err)
			}
			require.Equal(t, balance, testCase.balance)
		})
	}
}
