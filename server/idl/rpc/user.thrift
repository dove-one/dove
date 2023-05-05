namespace go user

include "../base/common.thrift"
include "../base/user.thrift"

struct LoginRequest {
    1:  string username (api.raw = "username" api.vd = "len($) > 0 && len($) < 33>")
    2:  string password (api.raw = "password" api.vd = "len($) > 0 && len($) < 33>")
}

struct LoginResponse {
    1:  common.BaseResponse base_resp
    2:  string token
}

service UserService {
    LoginResponse Login(1: LoginRequest req)

}