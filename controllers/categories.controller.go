package controllers

import (
	"codefood/models"
	"net/http"
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

func FindAllCategories(c echo.Context) error {
	limit := c.QueryParam("limit")
	skip := c.QueryParam("skip")
	countError := 0

	conv_limit, err := strconv.Atoi(limit)
	if err != nil {
		countError = countError + 1
	}
	conv_skip, err := strconv.Atoi(skip)
	if err != nil {
		countError = countError + 1
	}

	if countError == 0 {
		result, err := models.FindCategoriesAll(conv_limit, conv_skip)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
		}
		return c.JSON(http.StatusOK, result)
	}

	if countError > 1 {
		result, err := models.FindCategoriesAll2()
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
		}
		return c.JSON(http.StatusOK, result)
	}
	return c.JSON(http.StatusOK, nil)
}

func FindCategoriesById(c echo.Context) error {
	id := c.Param("id")

	conv_id, err := strconv.Atoi(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	result, err := models.FindCategoriesById(conv_id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}

func StoreCategories(c echo.Context) error {
	validate := validator.New()
	categories := &models.CategoriesValidation{}
	res := &models.Response{}

	err := c.Bind(categories)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	err = validate.Struct(categories)

	if err != nil {
		test := models.ResponseValidateError(err)
		res.Success = false
		res.Message = test.Error()
		res.Error = test
		return c.JSON(http.StatusBadRequest, res)
	}

	result, err := models.StoreCategories(categories.Name)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, result)
}

func UpdateCategories(c echo.Context) error {
	id := c.Param("id")
	validate := validator.New()
	categories := &models.CategoriesValidation{}
	res := &models.Response{}

	err := c.Bind(categories)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	err = validate.Struct(categories)
	if err != nil {
		test := models.ResponseValidateError(err)
		res.Success = false
		res.Message = test.Error()
		res.Error = test
		return c.JSON(400, res)
	}

	conv_id, err := strconv.Atoi(id) //convert to integer
	if err != nil {
		return c.JSON(400, err.Error())
	}

	errNumber, result, err := models.UpdateCategories(conv_id, categories.Name)
	if err != nil {
		return c.JSON(errNumber.Number, err.Error())
	}

	return c.JSON(errNumber.Number, result)
}

func DeleteCategories(c echo.Context) error {
	id := c.Param("id")

	conv_id, err := strconv.Atoi(id) //convert to integer
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	errNumber, result, err := models.DeleteCategories(conv_id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(errNumber.Number, result)
}
