package user

import "context"

type Storage interface {
	Create(ctx context.Context, user *CreateUserDto) (string, error)
	FindOne(ctx context.Context, id string) (u *User, err error)
	Update(ctx context.Context, user *User) error
	Delete(ctx context.Context, id int) error
}
