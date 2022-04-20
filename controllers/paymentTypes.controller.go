package controllers

import (
	"codefood/models"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

func GetJSONRawBody(c echo.Context) map[string]interface{} {

	jsonBody := make(map[string]interface{})
	err := json.NewDecoder(c.Request().Body).Decode(&jsonBody)
	if err != nil {
		return nil
	}

	return jsonBody
}

func FindAllPaymentTypes(c echo.Context) error {
	limit := c.QueryParam("limit")
	skip := c.QueryParam("skip")

	conv_limit, err := strconv.Atoi(limit)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	conv_skip, err := strconv.Atoi(skip)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	result, err := models.FindPaymentTypesAll(conv_limit, conv_skip)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}

func FindPaymentTypesById(c echo.Context) error {
	id := c.Param("id")

	conv_id, err := strconv.Atoi(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	result, err := models.FindPaymentTypesById(conv_id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}

func StorePaymentTypes(c echo.Context) error {
	validate := validator.New()
	paymentTypes := &models.PaymentTypesValidation{}
	res := &models.Response{}

	err := c.Bind(paymentTypes)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	err = validate.Struct(paymentTypes)

	if err != nil {
		if err != nil {
			test := models.ResponseValidateError(err)
			res.Success = false
			res.Message = test.Error()
			res.Error = test
			return c.JSON(http.StatusBadRequest, res)
		}
	}

	result, err := models.StorePaymentTypes(paymentTypes.Name, paymentTypes.Type, paymentTypes.Logo)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, result)
}

func UpdatePaymentTypes(c echo.Context) error {
	id := c.Param("id")
	validate := validator.New()
	paymentTypes := &models.PaymentTypesValidation{}
	res := &models.Response{}

	err := c.Bind(paymentTypes)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	err = validate.Struct(paymentTypes)
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

	errNumber, result, err := models.UpdatePaymentTypes(conv_id, paymentTypes.Name, paymentTypes.Type, paymentTypes.Logo)
	if err != nil {
		return c.JSON(errNumber.Number, err.Error())
	}

	return c.JSON(errNumber.Number, result)
}

func DeletePaymentTypes(c echo.Context) error {
	id := c.Param("id")

	conv_id, err := strconv.Atoi(id) //convert to integer
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	errNumber, result, err := models.DeletePaymentTypes(conv_id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(errNumber.Number, result)
}
