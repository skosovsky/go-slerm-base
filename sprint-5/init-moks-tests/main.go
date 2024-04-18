package main

import "fmt"

//go:generate mockgen -source=main.go -destination=mocks.go -package=main

type Command[T interface{}] struct {
	payload T
}

type IUseCase[T any, R any] interface {
	Handle(cmd Command[T]) (R, error)
}

type UserModel struct {
	ID   string
	Name string
}

type IUserRepository interface {
	FindUserById(id string) (*UserModel, error)
}

type INotificationService interface {
	NotifyUser(id string)
}

type FetchUserPayload struct {
	ID string
}

type FetchUserResult struct {
	us *UserModel
}

type FetchUsersUseCase struct {
	ur IUserRepository
	ns INotificationService
}

func (uc *FetchUsersUseCase) Handle(cmd Command[FetchUserPayload]) (*FetchUserResult, error) {
	useCase, err := uc.ur.FindUserById(cmd.payload.ID)
	if err != nil {
		return nil, fmt.Errorf("ndfind user by id filed: %w", err)
	}

	uc.ns.NotifyUser(useCase.ID)

	return &FetchUserResult{
		useCase,
	}, nil
}
