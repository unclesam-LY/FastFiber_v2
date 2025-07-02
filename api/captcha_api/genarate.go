package captcha_api

import (
	"FastFiber_v2/global"
	"FastFiber_v2/utils/captcha"
	"FastFiber_v2/utils/res"
	"github.com/gofiber/fiber/v2"
	"github.com/mojocn/base64Captcha"
	"go.uber.org/zap"
)

type GenerateResponse struct {
	CaptchaID string `json:"captcha_id"`
	Captcha   string `json:"captcha"`
}

func (CaptchaApi) GenerateView(c *fiber.Ctx) error {
	var driver = base64Captcha.DriverString{
		Height:          40,
		Width:           200,
		NoiseCount:      2,
		ShowLineOptions: 4,
		Length:          6,
		Source:          "0123456789",
	}

	cp := base64Captcha.NewCaptcha(&driver, captcha.CaptChaStore)
	id, b64s, _, err := cp.Generate()

	if err != nil {
		global.Log.Error("验证码生成失败", zap.Error(err))
		return res.FailWithMsg("验证码生成失败", c)
	}

	return res.OkWithData(GenerateResponse{
		CaptchaID: id,
		Captcha:   b64s,
	}, c)

}
