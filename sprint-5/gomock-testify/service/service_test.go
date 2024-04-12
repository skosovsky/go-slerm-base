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
				m.EXPECT().Info("send money error")
			},
			err: service.ErrConnection,
		},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			extDep := service.NewMockPaymentGateway(ctrl)
			if tc.initPaymentGateway != nil {
				tc.initPaymentGateway(extDep)
			}
			logger := service.NewMockLogger(ctrl)
			if tc.initLogger != nil {
				tc.initLogger(logger)
			}
			balance, err := service.TransferMoney(extDep, logger, tc.username, tc.amount)
			if err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.NoError(t, err)
			}
			require.Equal(t, balance, tc.balance)
		})
	}
}
