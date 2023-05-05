package main

import (
	"context"
	auth "github.com/dove-one/dove/server/common/kitex_gen/auth"
)

// AuthServiceImpl implements the last service interface defined in the IDL.
type AuthServiceImpl struct{}

// UserCheckAuth implements the AuthServiceImpl interface.
func (s *AuthServiceImpl) UserCheckAuth(ctx context.Context, req *auth.UserCheckAuthReq) (resp *auth.UserCheckAuthResp, err error) {
	// TODO: Your code here...

	return
}
