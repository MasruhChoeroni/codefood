package controllers

import (
	"codefood/models"
	"encoding/json"
	"net/http"
	"strconv"

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
	name := c.FormValue("name")
	tipe := c.FormValue("type")
	logo := c.FormValue("logo")

	// validate = validator.New()

	// type PaymentTypesValidate struct {
	// 	Name string `validate:"required"`
	// 	Type string `validate:"required"`
	// }

	// outer := &PaymentTypesValidate{
	// 	Name: name,
	// 	Type: tipe,
	// }

	// err := validate.Struct(outer)

	// if err != nil {
	// 	report, ok := err.(*echo.HTTPError)
	// 	if !ok {
	// 		report = echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	// 	}

	// 	if castedObject, ok := err.(validator.ValidationErrors); ok {
	// 		for _, err := range castedObject {
	// 			switch err.Tag() {
	// 			case "required":
	// 				report.Message = fmt.Sprintf("%s is required",
	// 					err.Field())
	// 			case "len":
	// 				report.Message = fmt.Sprintf("%s value length must be %s",
	// 					err.Field(), err.Param())
	// 			case "numeric":
	// 				report.Message = fmt.Sprintf("%s value must be numeric",
	// 					err.Field())
	// 			}
	// 		}
	// 	}
	// 	return c.JSON(http.StatusInternalServerError, report)
	// }

	result, err := models.StorePaymentTypes(name, tipe, logo)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, result)
}

func UpdatePaymentTypes(c echo.Context) error {
	id := c.Param("id")
	name := c.FormValue("name")
	tipe := c.FormValue("type")
	logo := c.FormValue("logo")

	conv_id, err := strconv.Atoi(id) //convert to integer
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	result, err := models.UpdatePaymentTypes(conv_id, name, tipe, logo)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, result)
}

func DeletePaymentTypes(c echo.Context) error {
	id := c.Param("id")

	conv_id, err := strconv.Atoi(id) //convert to integer
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	result, err := models.DeletePaymentTypes(conv_id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, result)
}
