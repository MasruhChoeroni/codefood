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

func FindCashiersAll(limit int, skip int) (Response, error) {
	var obj Cashiers
	var arrobj []Cashiers
	var res Response

	con := db.CreateCon()

	sqlStatement := "SELECT * FROM cashiers LIMIT ?, ?"

	rows, err := con.Query(sqlStatement, limit, skip)

	defer rows.Close()

	if err != nil {
		return res, err
	}

	for rows.Next() {
		err = rows.Scan(&obj.Id, &obj.Name, &obj.Passcode, &obj.CreatedAt, &obj.UpdatedAt)
		if err != nil {
			return res, err
		}
		arrobj = append(arrobj, obj)
	}

	res.Success = true
	res.Message = "Success"
	res.Data = arrobj

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
