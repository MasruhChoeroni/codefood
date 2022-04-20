package controllers

import (
	"codefood/models"
	"net/http"
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

func FindAllCashiers(c echo.Context) error {
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
		result, err := models.FindCashiersAll(conv_limit, conv_skip)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
		}
		return c.JSON(http.StatusOK, result)
	}

	if countError > 1 {
		result, err := models.FindCashiersAll2()
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
		}
		return c.JSON(http.StatusOK, result)
	}
	return c.JSON(http.StatusOK, nil)

}

func FindCashiersById(c echo.Context) error {
	id := c.Param("id")

	conv_id, err := strconv.Atoi(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	result, err := models.FindCashiersById(conv_id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}

func FindCashiersPasscodeById(c echo.Context) error {
	id := c.Param("id")

	conv_id, err := strconv.Atoi(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	httpNumber, result, err := models.FindCashiersPasscodeById(conv_id)
	if err != nil {
		return c.JSON(httpNumber.Number, map[string]string{"message": err.Error()})
	}

	return c.JSON(httpNumber.Number, result)
}

func StoreCashiers(c echo.Context) error {
	validate := validator.New()
	cashier := &models.CashiersValidation{}
	res := &models.Response{}

	err := c.Bind(cashier)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	err = validate.Struct(cashier)

	if err != nil {
		test := models.ResponseValidateError(err)
		res.Success = false
		res.Message = test.Error()
		res.Error = test
		return c.JSON(http.StatusBadRequest, res)
	}

	result, err := models.StoreCashiers(cashier.Name, cashier.Passcode)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, result)
}

func UpdateCashiers(c echo.Context) error {
	id := c.Param("id")
	validate := validator.New()
	cashier := &models.CashiersValidation{}
	res := &models.Response{}

	err := c.Bind(cashier)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	err = validate.Struct(cashier)

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

	errNumber, result, err := models.UpdateCashiers(conv_id, cashier.Name, cashier.Passcode)
	if err != nil {
		return c.JSON(errNumber.Number, err.Error())
	}

	return c.JSON(errNumber.Number, result)
}

func DeleteCashiers(c echo.Context) error {
	id := c.Param("id")

	conv_id, err := strconv.Atoi(id) //convert to integer
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	errNumber, result, err := models.DeleteCashiers(conv_id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(errNumber.Number, result)
}
