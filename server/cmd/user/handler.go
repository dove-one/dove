package main

import (
	"context"
	"fmt"
	u "github.com/dove-one/dove/server/cmd/user/model/user"
	user "github.com/dove-one/dove/server/common/kitex_gen/user"
)

// UserServiceImpl implements the last service interface defined in the IDL.
type UserServiceImpl struct {
	u.MysqlManagerUser
}

// Login implements the UserServiceImpl interface.
func (s *UserServiceImpl) Login(ctx context.Context, req *user.LoginRequest) (resp *user.LoginResponse, err error) {
	// TODO: Your code here...
	condition, err := s.MysqlManagerUser.GetUsersByCondition(ctx, u.Uuid(1))
	if err != nil {
		return nil, err
	}
	fmt.Println(condition[0].Account)
	return
}
