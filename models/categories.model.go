package models

import (
	"codefood/db"
	"time"
)

type Categories struct {
	Id        int       `json:"categoryId"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"updatedAt"`
	UpdatedAt time.Time `json:"createdAt"`
}

type CategoriesValidation struct {
	Name string `json:"name" validate:"required"`
}

type Categories2 struct {
	Id   int    `json:"categoryId"`
	Name string `json:"name"`
}

func FindCategoriesAll(limit int, skip int) (Response, error) {
	var obj Categories2
	var arrobj []Categories2
	var res Response
	var total int64

	con := db.CreateCon()

	sqlCountStatement := "SELECT COUNT(*) AS total FROM categories"

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

	sqlStatement := "SELECT id, name FROM categories LIMIT ? OFFSET ?"

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
		"categories": arrobj,
		"meta": map[string]int64{
			"total": total,
			"limit": int64(limit),
			"skip":  int64(skip),
		},
	}

	return res, nil
}

func FindCategoriesById(id int) (Response, error) {
	var obj Categories2
	var res Response

	con := db.CreateCon()

	sqlStatement := "SELECT id, name FROM categories WHERE id = ?"

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

func StoreCategories(name string) (Response, error) {
	var res Response
	var obj Categories

	con := db.CreateCon()

	sqlStatement := "INSERT categories (name, created_at, updated_at) VALUES (?, now(), now());"

	stmt, err := con.Prepare(sqlStatement)
	if err != nil {
		return res, err
	}

	result, err := stmt.Exec(name)
	if err != nil {
		return res, err
	}

	lastInsertedId, err := result.LastInsertId()
	if err != nil {
		return res, err
	}

	sqlStatement2 := "SELECT * FROM categories WHERE id = ?;"
	rows, err := con.Query(sqlStatement2, lastInsertedId)
	if err != nil {
		return res, err
	}

	defer rows.Close()
	for rows.Next() {
		err = rows.Scan(&obj.Id, &obj.Name, &obj.CreatedAt, &obj.UpdatedAt)
		if err != nil {
			return res, err
		}
	}

	res.Success = true
	res.Message = "Success"
	res.Data = obj

	return res, nil
}

func UpdateCategories(id int, name string) (Response, error) {
	var res Response

	con := db.CreateCon()

	sqlStatement := "UPDATE categories SET name= ?, updated_at = now() WHERE id = ?;"

	stmt, err := con.Prepare(sqlStatement)
	if err != nil {
		return res, err
	}

	result, err := stmt.Exec(name, id)
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

func DeleteCategories(id int) (Response, error) {
	var res Response

	con := db.CreateCon()

	sqlStatement := "DELETE FROM categories WHERE id = ?;"

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
