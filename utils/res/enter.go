package res

import (
	"FastFiber_v2/utils/validata"
	"github.com/gofiber/fiber/v2"
)

type Response struct {
	Code int    `json:"code"`
	Data any    `json:"data"`
	Msg  string `json:"msg"`
}

const (
	Success = 0
	Error   = 7
)

// 成功响应

func Ok(code int, data any, msg string, c *fiber.Ctx) error {
	return c.JSON(Response{
		Code: code,
		Data: data,
		Msg:  msg,
	})
}

func OkWithMsg(msg string, c *fiber.Ctx) error {
	return Ok(Success, struct{}{}, msg, c)
}

func OkWithList(list any, count int64, c *fiber.Ctx) error {
	return Ok(
		Success,
		map[string]any{
			"list":  list,
			"count": count,
		}, "成功", c)
}

func OkWithData(data any, c *fiber.Ctx) error {
	return Ok(Success, data, "成功", c)
}

// 失败响应

func Fail(code int, msg string, c *fiber.Ctx) error {
	return c.JSON(Response{
		Code: code,
		Data: struct{}{},
		Msg:  msg,
	})
}

func FailWithMsg(msg string, c *fiber.Ctx) error {
	return Fail(Error, msg, c)
}

func FailWithError(err error, c *fiber.Ctx) error {
	msg := validata.ValidateErr(err)
	return Fail(Error, msg, c)
}
