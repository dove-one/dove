namespace go auth

include "../base/common.thrift"

struct UserCheckAuthResp{
    1: common.BaseResponse base_resp
    2: string uuid
}

struct UserCheckAuthReq{
    1: string token
}


service AuthService{
  UserCheckAuthResp UserCheckAuth(1: UserCheckAuthReq req)

}