package models

import (
	"codefood/db"
	"time"
)

type Cashiers struct {
	Id        int       `json:"cashierId"`
	Name      string    `json:"name"`
	Passcode  string    `json:"passcode"`
	CreatedAt time.Time `json:"updatedAt"`
	UpdatedAt time.Time `json:"createdAt"`
}

type CashiersPostValidation struct {
	Name     string `json:"name" validate:"required"`
	Passcode string `json:"passcode" validate:"required,numeric,len=6"`
}

type Cashiers2 struct {
	Id   int    `json:"cashierId"`
	Name string `json:"name"`
}

func FindCashiersAll(limit int, skip int) (Response, error) {
	var obj Cashiers2
	var arrobj []Cashiers2
	var res Response
	var total int64

	con := db.CreateCon()

	sqlCountStatement := "SELECT COUNT(*) AS total FROM cashiers"

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

	sqlStatement := "SELECT id, name FROM cashiers LIMIT ? OFFSET ?"

	rows, err := con.Query(sqlStatement, limit, skip)
	if err != nil {
		return res, err
	}
	defer rows.Close()

	for rows.Next() {
		err = rows.Scan(&obj.Id, &obj.Name)
		if err != nil {
			return res, err
		}
		arrobj = append(arrobj, obj)
	}

	res.Success = true
	res.Message = "Success"
	res.Data = map[string]interface{}{
		"cashiers": arrobj,
		"meta": map[string]int64{
			"total": total,
			"limit": int64(limit),
			"skip":  int64(skip),
		},
	}

	return res, nil
}

func FindCashiersById(id int) (Response, error) {
	var obj Cashiers2
	var res Response

	con := db.CreateCon()

	sqlStatement := "SELECT id, name FROM cashiers WHERE id = ?"

	rows, err := con.Query(sqlStatement, id)
	if err != nil {
		return res, err
	}
	defer rows.Close()

	for rows.Next() {
		err = rows.Scan(&obj.Id, &obj.Name)
		if err != nil {
			return res, err
		}
	}

	res.Success = true
	res.Message = "Success"
	res.Data = obj

	return res, nil
}

func FindCashiersPasscodeById(id int) (Response, error) {
	var res Response
	var passcode string

	con := db.CreateCon()

	sqlStatement := "SELECT passcode FROM cashiers WHERE id = ?"

	rows, err := con.Query(sqlStatement, id)
	if err != nil {
		return res, err
	}
	defer rows.Close()

	for rows.Next() {
		err = rows.Scan(&passcode)
		if err != nil {
			return res, err
		}
	}

	res.Success = true
	res.Message = "Success"
	res.Data = map[string]string{
		"passcode": passcode,
	}

	return res, nil
}

func StoreCashiers(name string, passcode string) (Response, error) {
	var res Response
	var obj Cashiers

	con := db.CreateCon()

	sqlStatement := "INSERT cashiers (name, passcode, created_at, updated_at) VALUES (?, ?, now(), now());"

	stmt, err := con.Prepare(sqlStatement)
	if err != nil {
		return res, err
	}

	result, err := stmt.Exec(name, passcode)
	if err != nil {
		return res, err
	}

	lastInsertedId, err := result.LastInsertId()
	if err != nil {
		return res, err
	}

	sqlStatement2 := "SELECT * FROM cashiers WHERE id = ?;"
	rows, err := con.Query(sqlStatement2, lastInsertedId)

	defer rows.Close()

	if err != nil {
		return res, err
	}

	for rows.Next() {
		err = rows.Scan(&obj.Id, &obj.Name, &obj.Passcode, &obj.CreatedAt, &obj.UpdatedAt)
		if err != nil {
			return res, err
		}
	}

	res.Success = true
	res.Message = "Success"
	res.Data = obj

	return res, nil
}

func UpdateCashiers(id int, name string, passcode string) (Response, error) {
	var res Response

	con := db.CreateCon()

	sqlStatement := "UPDATE cashiers SET name= ?, passcode = ?, updated_at = now() WHERE id = ?;"

	stmt, err := con.Prepare(sqlStatement)
	if err != nil {
		return res, err
	}

	result, err := stmt.Exec(name, passcode, id)
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

func DeleteCashiers(id int) (Response, error) {
	var res Response

	con := db.CreateCon()

	sqlStatement := "DELETE FROM cashiers WHERE id = ?;"

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
