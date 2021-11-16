package controllers

import (
	"e-commerce-transaction/services"

	"github.com/gofiber/fiber/v2"
)

type SellersController struct{}

func (s SellersController) GetAllSellers(ctx *fiber.Ctx) error {
	sellers, err := services.GetAllSellers()

	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(&fiber.Map{
			"message": err.Error(),
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(sellers)
}

func (s SellersController) GetTotalSellsAggregateByState(ctx *fiber.Ctx) error {
	sellers, err := services.GetTotalSellsAggregateByState()

	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(&fiber.Map{
			"message": err.Error(),
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(sellers)
}

func (s SellersController) GetSellersByState(ctx *fiber.Ctx) error {
	state := ctx.Params("state")
	sellers, err := services.GetSellersByState(state)

	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(&fiber.Map{
			"message": err.Error(),
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(sellers)
}
