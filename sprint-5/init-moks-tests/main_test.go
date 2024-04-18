package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestUseCase_Handle(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockUser := UserModel{
		ID:   "1",
		Name: "",
	}

	mockedUserRepo := NewMockIUserRepository(ctrl)
	mockedNotificationService := NewMockINotificationService(ctrl)

	mockedUserRepo.EXPECT().FindUserById(mockUser.ID).Return(&mockUser, nil)
	mockedNotificationService.EXPECT().NotifyUser(mockUser.ID).MaxTimes(1)

	useCase := &FetchUsersUseCase{
		ur: mockedUserRepo,
		ns: mockedNotificationService,
	}

	cmd := Command[FetchUserPayload]{
		payload: FetchUserPayload{
			ID: "1",
		},
	}

	result, err := useCase.Handle(cmd)
	if err != nil {
		return
	}

	assert.Equal(t, result.us.ID, cmd.payload.ID)
}
