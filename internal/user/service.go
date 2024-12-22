package user

import (
	"context"
	"rest/pkg/logging"
)

type service struct {
	storage Storage
	logger  *logging.Logger
}

func (s *service) Create(ctx context.Context, dto *CreateUserDto) (u User, err error) {
	//TODO next
	return
}
