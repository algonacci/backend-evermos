package handler

import (
	"strconv"

	"github.com/algonacci/backend-evermos/model"
	"github.com/algonacci/backend-evermos/repository"
	"github.com/gofiber/fiber/v2"
)

func GetShop(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return err
	}

	shop, err := repository.FindShopByID(uint(id))
	if err != nil {
		return err
	}

	return c.JSON(shop)
}

func CreateShop(c *fiber.Ctx) error {
	req := new(model.Shop)
	if err := c.BodyParser(req); err != nil {
		return err
	}

	shop, err := repository.CreateShop(req)
	if err != nil {
		return err
	}

	return c.JSON(shop)
}

func UpdateShop(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return err
	}

	req := new(model.Shop)
	if err := c.BodyParser(req); err != nil {
		return err
	}

	shop, err := repository.UpdateShop(uint(id), req)
	if err != nil {
		return err
	}

	return c.JSON(shop)
}

func DeleteShop(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return err
	}

	err = repository.DeleteShop(uint(id))
	if err != nil {
		return err
	}

	return c.SendString("Shop deleted successfully")
}
