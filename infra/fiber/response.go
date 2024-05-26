package fiber

import (
	"e-commerce-gorm/infra/response"
	"github.com/gofiber/fiber/v2"
)

type Response struct {
	HttpCode  int         `json:"http_code"`
	Success   bool        `json:"success"`
	Message   string      `json:"message"`
	Payload   interface{} `json:"payload,omitempty"`
	Query     interface{} `json:"query,omitempty"`
	Error     string      `json:"error,omitempty"`
	ErrorCode string      `json:"errorCode,omitempty"`
}

func NewResponse(params ...func(*Response) *Response) Response {
	var resp = Response{
		Success: true,
	}

	for _, param := range params {
		param(&resp)
	}
	return resp
}

func WithHttpCode(httpCode int) func(*Response) *Response {
	return func(response *Response) *Response {
		response.HttpCode = httpCode
		return response
	}
}

func WithMessage(message string) func(*Response) *Response {
	return func(response *Response) *Response {
		response.Message = message
		return response
	}
}

func WithPayload(payload interface{}) func(*Response) *Response {
	return func(response *Response) *Response {
		response.Payload = payload
		return response
	}
}

func WithQuery(query int) func(*Response) *Response {
	return func(response *Response) *Response {
		response.Query = query
		return response
	}
}

func WithError(err error) func(*Response) *Response {
	return func(r *Response) *Response {
		r.Success = false

		myErr, ok := err.(response.Error)
		if !ok {
			myErr = response.ErrorGeneral
		}

		r.Error = myErr.Message
		r.ErrorCode = myErr.Code
		r.HttpCode = myErr.HttpCode

		return r
	}
}

func (respond Response) Send(ctx *fiber.Ctx) error {
	return ctx.Status(respond.HttpCode).JSON(respond)
}
