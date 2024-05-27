package product

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"
)

func Init(router fiber.Router, db *sqlx.DB) {
	repo := newRepository(db)
	svc := newService(repo)
	handler := newHandler(svc)

	productRoute := router.Group("products")
	{
		productRoute.Post("", handler.CreateProduct)
		productRoute.Get("", handler.GetListProduct)
		productRoute.Get("/sku/:sku", handler.GetProductDetail)
	}
}
