package models

import (
	"codefood/db"
	"time"
)

type PaymentTypes struct {
	Id        int       `json:"paymentId"`
	Name      string    `json:"name"`
	Type      string    `json:"type"`
	Logo      string    `json:"logo"`
	CreatedAt time.Time `json:"updatedAt"`
	UpdatedAt time.Time `json:"createdAt"`
}

type PaymentTypesValidation struct {
	Name string `json:"name" validate:"required"`
	Type string `json:"type" validate:"required"`
	Logo string `json:"logo"`
}

type PaymentTypes2 struct {
	Id   int    `json:"paymentId"`
	Name string `json:"name"`
	Type string `json:"type"`
	Logo string `json:"logo"`
}

func FindPaymentTypesAll(limit int, skip int) (Response, error) {
	var obj PaymentTypes2
	var arrobj []PaymentTypes2
	var res Response
	var total int64

	con := db.CreateCon()

	sqlCountStatement := "SELECT COUNT(*) AS total FROM payment_types"

	row, err := con.Query(sqlCountStatement)
	if err != nil {
		return res, err
	}
	defer row.Close()
	for row.Next() {
		err = row.Scan(&total)
		if err != nil {
			return res, err
		}
	}

	sqlStatement := "SELECT id, name, type, logo FROM payment_types LIMIT ? OFFSET ?"

	rows, err := con.Query(sqlStatement, limit, skip)
	if err != nil {
		return res, err
	}
	defer rows.Close()

	for rows.Next() {
		err = rows.Scan(&obj.Id, &obj.Name, &obj.Type, &obj.Logo)
		if err != nil {
			return res, err
		}
		arrobj = append(arrobj, obj)
	}

	res.Success = true
	res.Message = "Success"
	res.Data = map[string]interface{}{
		"payments": arrobj,
		"meta": map[string]int64{
			"total": total,
			"limit": int64(limit),
			"skip":  int64(skip),
		},
	}

	return res, nil
}

func FindPaymentTypesById(id int) (Response, error) {
	var obj PaymentTypes2
	var res Response

	con := db.CreateCon()

	sqlStatement := "SELECT id, name, type, logo FROM payment_types WHERE id = ?"

	rows, err := con.Query(sqlStatement, id)
	if err != nil {
		return res, err
	}
	defer rows.Close()

	for rows.Next() {
		err = rows.Scan(&obj.Id, &obj.Name, &obj.Type, &obj.Logo)
		if err != nil {
			return res, err
		}
	}

	res.Success = true
	res.Message = "Success"
	res.Data = obj

	return res, nil
}

func StorePaymentTypes(name string, tipe string, logo string) (Response, error) {
	var res Response
	var obj PaymentTypes

	con := db.CreateCon()

	sqlStatement := "INSERT payment_types (name, type, logo, created_at, updated_at) VALUES (?, ?, ?, now(), now());"

	stmt, err := con.Prepare(sqlStatement)
	if err != nil {
		return res, err
	}

	result, err := stmt.Exec(name, tipe, logo)
	if err != nil {
		return res, err
	}

	lastInsertedId, err := result.LastInsertId()
	if err != nil {
		return res, err
	}

	sqlStatement2 := "SELECT * FROM payment_types WHERE id = ?;"
	rows, err := con.Query(sqlStatement2, lastInsertedId)

	defer rows.Close()

	if err != nil {
		return res, err
	}

	for rows.Next() {
		err = rows.Scan(&obj.Id, &obj.Name, &obj.Type, &obj.Logo, &obj.CreatedAt, &obj.UpdatedAt)
		if err != nil {
			return res, err
		}
	}

	res.Success = true
	res.Message = "Success"
	res.Data = obj

	return res, nil
}

func UpdatePaymentTypes(id int, name string, tipe string, logo string) (Response, error) {
	var res Response

	con := db.CreateCon()

	sqlStatement := "UPDATE payment_types SET name= ?, type= ?, logo= ?, updated_at = now() WHERE id = ?;"

	stmt, err := con.Prepare(sqlStatement)
	if err != nil {
		return res, err
	}

	result, err := stmt.Exec(name, tipe, logo, id)
	if err != nil {
		return res, err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return res, err
	}

	res.Success = true
	res.Message = "Success"
	res.Data = map[string]int64{
		"rows_affected": rowsAffected,
	}

	return res, nil
}

func DeletePaymentTypes(id int) (Response, error) {
	var res Response

	con := db.CreateCon()

	sqlStatement := "DELETE FROM payment_types WHERE id = ?;"

	stmt, err := con.Prepare(sqlStatement)
	if err != nil {
		return res, err
	}

	result, err := stmt.Exec(id)
	if err != nil {
		return res, err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return res, err
	}

	res.Success = true
	res.Message = "Success"
	res.Data = map[string]int64{
		"rows_affected": rowsAffected,
	}

	return res, nil
}
