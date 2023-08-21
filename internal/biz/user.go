package biz

import (
	"context"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	Username     string
	Email        string
	Token        string
	Bio          string
	Image        string
	Password     string
	PasswordHash string
}

type UserRepo interface {
	// CreateUser 这里需要把领域对象传进去
	CreateUser(ctx context.Context, user *User) error
	GetUserByEmail(ctx context.Context, email string) (*User, error)
}

type UserUseCase struct {
	ur UserRepo
	pr ProfileRepo
}

func NewUserUseCase(ur UserRepo, pr ProfileRepo) *UserUseCase {
	return &UserUseCase{ur: ur, pr: pr}
}

func (uc *UserUseCase) Register(ctx context.Context, user *User) error {
	return uc.ur.CreateUser(ctx, user)
}

func (uc *UserUseCase) Login(ctx context.Context, email, password string) (*User, error) {
	user, err := uc.ur.GetUserByEmail(ctx, email)
	if err != nil {
		return nil, err
	}
	return user, nil
}

// 1.定义data层的接口方法
// 2.定义领域对象
// 3.通过data层的接口方法去操作领域对象

type ProfileRepo interface {
}

func hashPassword(pwd string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(pwd), bcrypt.DefaultCost)
}
