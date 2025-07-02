package api

import (
	"FastFiber_v2/api/captcha_api"
	"FastFiber_v2/api/user_api"
)

type Api struct {
	UserApi    user_api.UserApi
	CaptchaApi captcha_api.CaptchaApi
}

var App = new(Api)
