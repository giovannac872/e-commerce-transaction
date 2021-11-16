package controllers

import (
	"e-commerce-transaction/services"

	"github.com/gofiber/fiber/v2"
)

type ProductsController struct{}

func (p ProductsController) GetAllProducts(ctx *fiber.Ctx) error {
	products, err := services.GetAllProducts()

	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(&fiber.Map{
			"message": err.Error(),
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(products)
}

func (p ProductsController) GetProductsByCategory(ctx *fiber.Ctx) error {
	slugCategory := ctx.Params("category")
	products, err := services.GetAllProductsByCategory(slugCategory)

	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(&fiber.Map{
			"message": err.Error(),
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(products)

}

func (p ProductsController) GetMeanProductsPriceAggregateByCategory(ctx *fiber.Ctx) error {
	products, err := services.GetMeanProductsPriceAggregateByCategory()

	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(&fiber.Map{
			"message": err.Error(),
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(products)
}
