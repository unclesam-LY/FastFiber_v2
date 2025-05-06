package api

import "FastFiber_v2/api/user_api"

type Api struct {
	UserApi user_api.UserApi
}

var App = new(Api)
