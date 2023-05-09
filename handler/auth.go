package handler

import (
	"github.com/algonacci/backend-evermos/config"
	"github.com/algonacci/backend-evermos/model"
	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
)

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginResponse struct {
	Token string `json:"token"`
}

type RegisterRequest struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Phone    string `json:"phone"`
	Role     string `json:"role"`
}

type RegisterResponse struct {
	Token string `json:"token"`
}

func Login(c *fiber.Ctx) error {
	req := new(LoginRequest)
	if err := c.BodyParser(req); err != nil {
		return err
	}

	// Check user exists
	db, err := config.ConnectDB()
	if err != nil {
		return err
	}
	defer func() {
		sqlDB, err := db.DB()
		if err != nil {
			return
		}
		sqlDB.Close()
	}()

	var user model.User
	if err := db.Where("email = ?", req.Email).First(&user).Error; err != nil {
		return fiber.ErrUnauthorized
	}

	// Check password
	if user.Password != req.Password {
		return fiber.ErrUnauthorized
	}

	// Generate JWT token
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["id"] = user.ID
	claims["email"] = user.Email
	tokenString, err := token.SignedString([]byte("secret"))
	if err != nil {
		return err
	}

	// Return token
	resp := LoginResponse{Token: tokenString}
	return c.JSON(resp)
}

func Register(c *fiber.Ctx) error {
	req := new(RegisterRequest)
	if err := c.BodyParser(req); err != nil {
		return err
	}

	// Check existing user
	db, err := config.ConnectDB()
	if err != nil {
		return err
	}
	defer func() {
		sqlDB, err := db.DB()
		if err != nil {
			return
		}
		sqlDB.Close()
	}()
	var user model.User
	if db.Where("email = ?", req.Email).First(&user).Error == nil {
		return fiber.ErrBadRequest
	}
	if db.Where("phone = ?", req.Phone).First(&user).Error == nil {
		return fiber.ErrBadRequest
	}

	// Create user
	user = model.User{
		Name:     req.Name,
		Email:    req.Email,
		Password: req.Password,
		Phone:    req.Phone,
		Role:     req.Role,
	}
	if err := db.Create(&user).Error; err != nil {
		return err
	}

	// Create shop
	shop := model.Shop{Name: req.Name, UserID: user.ID}
	if err := db.Create(&shop).Error; err != nil {
		return err
	}
	user.ShopID = shop.ID
	if err := db.Save(&user).Error; err != nil {
		return err
	}

	// Generate JWT token
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["id"] = user.ID
	claims["email"] = user.Email
	tokenString, err := token.SignedString([]byte("secret"))
	if err != nil {
		return err
	}

	// Return token
	resp := RegisterResponse{Token: tokenString}
	return c.JSON(resp)
}
