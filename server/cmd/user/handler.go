package main

import (
	"context"
	user "github.com/dove-one/dove/server/common/kitex_gen/user"
)

// UserServiceImpl implements the last service interface defined in the IDL.
type UserServiceImpl struct {
	EncryptManager
}

type EncryptManager interface {
	EncryptPassword(code string) string
}

// Login implements the UserServiceImpl interface.
func (s *UserServiceImpl) Login(ctx context.Context, req *user.LoginRequest) (resp *user.LoginResponse, err error) {
	// TODO: Your code here...
	return
}
