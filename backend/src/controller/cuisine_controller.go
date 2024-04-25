package controller

import (
	"backend/src/model"
	"backend/src/usecase"
	"net/http"
	"strconv"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

type ICuisineController interface {
	GetAllCuisines(c echo.Context) error
	GetCuisineById(c echo.Context) error
	CreateCuisine(c echo.Context) error
	UpdateCuisine(c echo.Context) error
	DeleteCuisine(c echo.Context) error
}

type cuisineController struct {
	cu usecase.ICuisineUsecase
}

func NewCuisineController(cu usecase.ICuisineUsecase) ICuisineController {
	return &cuisineController{cu}
}

func (cc *cuisineController) GetAllCuisines(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)    //コンテキストからjwtをデコードした値を読み込む
	claims := user.Claims.(jwt.MapClaims) //その中のデコードされたclaimsを取得
	userId := claims["user_id"]           //claimsの中のuserIdを取得

	cuisineRes, err := cc.cu.GetAllCuisines(uint(userId.(float64))) //一度floatにしてからuintに型変換
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, cuisineRes)
}

func (cc *cuisineController) GetCuisineById(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	userId := claims["user_id"]
	id := c.Param("cuisineId")       //リクエストパラメーターからcuisineIdを取得
	cuisineId, _ := strconv.Atoi(id) //stringからintに
	cuisineRes, err := cc.cu.GetCuisineById(uint(userId.(float64)), uint(cuisineId))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, cuisineRes)
}

func (cc *cuisineController) CreateCuisine(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	userId := claims["user_id"]

	cuisine := model.Cuisine{}
	if err := c.Bind(&cuisine); err != nil { //リクエストボディに含まれる内容をcuisine構造体に代入
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	cuisine.UserId = uint(userId.(float64))
	cuisineRes, err := cc.cu.CreateCuisine(cuisine)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusCreated, cuisineRes)
}

func (cc *cuisineController) UpdateCuisine(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	userId := claims["user_id"]
	id := c.Param("cuisineId")
	cuisineId, _ := strconv.Atoi(id)

	cuisine := model.Cuisine{}
	if err := c.Bind(&cuisine); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	cuisineRes, err := cc.cu.UpdateCuisine(cuisine, uint(userId.(float64)), uint(cuisineId))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, cuisineRes)
}

func (cc *cuisineController) DeleteCuisine(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	userId := claims["user_id"]
	id := c.Param("cuisineId")
	cuisineId, _ := strconv.Atoi(id)

	cuisine := model.Cuisine{}
	if err := c.Bind(&cuisine); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	err := cc.cu.DeleteCuisine(uint(userId.(float64)), uint(cuisineId))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.NoContent(http.StatusNoContent)
}

// func (cc * cuisineController) UploadImage(c echo.Context) error{
// 	user := c.Get("user").(*jwt.Token)
// 	claims := user.Claims.(jwt.MapClaims)
// 	userId := claims["user_id"]

// 	file,err := c.FormFile("image")
// }
