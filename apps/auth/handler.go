package auth

import (
	infrafiber "e-commerce-gorm/infra/fiber"
	"e-commerce-gorm/infra/response"
	"github.com/gofiber/fiber/v2"
	"net/http"
)

type handler struct {
	svc service
}

func newHandler(svc service) handler {
	return handler{
		svc: svc,
	}
}

func (h handler) register(ctx *fiber.Ctx) error {
	var req RegisterRequestPayload // Pastikan tipe data sesuai untuk register

	if err := ctx.BodyParser(&req); err != nil {
		myErr := response.ErrorBadRequest
		return infrafiber.NewResponse(
			infrafiber.WithMessage(err.Error()),
			infrafiber.WithError(myErr),
			infrafiber.WithHttpCode(http.StatusBadRequest),
			infrafiber.WithMessage("register fail"), // Perbaiki typo dari "regitser fail" ke "register fail"
		).Send(ctx)
	}

	if err := h.svc.register(ctx.UserContext(), req); err != nil {
		myErr, ok := response.ErrorMapping[err.Error()]
		if !ok {
			myErr = response.ErrorGeneral
		}

		return infrafiber.NewResponse(
			infrafiber.WithMessage(err.Error()),
			infrafiber.WithError(myErr)).Send(ctx)
	}
	return infrafiber.NewResponse(
		infrafiber.WithHttpCode(http.StatusCreated),
		infrafiber.WithMessage("register success"), // Perbaiki message dari "login success" ke "register success"
	).Send(ctx)
}

func (h handler) login(ctx *fiber.Ctx) error {
	var req LoginRequestPayload // Pastikan tipe data sesuai untuk login

	if err := ctx.BodyParser(&req); err != nil {
		myErr := response.ErrorBadRequest
		return infrafiber.NewResponse(
			infrafiber.WithMessage(err.Error()),
			infrafiber.WithError(myErr),
			infrafiber.WithHttpCode(http.StatusBadRequest),
			infrafiber.WithMessage("login fail"),
		).Send(ctx)
	}

	token, err := h.svc.login(ctx.UserContext(), req)
	if err != nil {
		myErr, ok := response.ErrorMapping[err.Error()]
		if !ok {
			myErr = response.ErrorGeneral
		}

		return infrafiber.NewResponse(
			infrafiber.WithMessage(err.Error()),
			infrafiber.WithError(myErr),
		).Send(ctx)
	}
	return infrafiber.NewResponse(
		infrafiber.WithHttpCode(http.StatusOK), // Ganti ke http.StatusOK untuk login success
		infrafiber.WithPayload(map[string]interface{}{
			"access_token": token,
		}),
		infrafiber.WithMessage("login success"),
	).Send(ctx)
}
