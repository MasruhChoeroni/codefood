package models

import (
	"codefood/db"
	"time"
)

type Products struct {
	Id         int        `json:"paymentId"`
	Name       string     `json:"name"`
	CategoryId string     `json:"type"`
	Sku        string     `json:"sku"`
	Stock      int64      `json:"stock"`
	Price      int64      `json:"price"`
	DiscountId int64      `json:"discountId"`
	Discount   []Discount `json:"discount"`
	Image      string     `json:"image"`
	CreatedAt  time.Time  `json:"updatedAt"`
	UpdatedAt  time.Time  `json:"createdAt"`
}

type Products2 struct {
	Id         int        `json:"paymentId"`
	Name       string     `json:"name"`
	CategoryId string     `json:"type"`
	Sku        string     `json:"sku"`
	Stock      int64      `json:"stock"`
	Price      int64      `json:"price"`
	DiscountId int64      `json:"discountId"`
	Discount   []Discount `json:"discount"`
	Image      string     `json:"image"`
}

type Discount struct {
	Id        int    `json:"discountId"`
	Qty       int8   `json:"qty"`
	Type      string `json:"type"`
	Result    int64  `json:"result"`
	ExpiredAt int64  `json:"expiredAt"`
}

func FindProductsAll(limit int, skip int, categoryId int, searchName string) (Response, error) {
	var obj Products2
	var arrobj []Products2
	var res Response
	var total int64

	con := db.CreateCon()

	sqlCountStatement := "SELECT COUNT(*) AS total FROM products WHERE category_id = ? AND name LIKE '%?%'"

	row, err := con.Query(sqlCountStatement, categoryId, searchName)
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

	sqlStatement := "SELECT id, name, sku, stock, price, discount, image, category_id FROM products WHERE category_id = ? AND name LIKE '%?%' LIMIT ? OFFSET ?"

	rows, err := con.Query(sqlStatement, categoryId, searchName, limit, skip)
	if err != nil {
		return res, err
	}
	defer rows.Close()

	for rows.Next() {
		err = rows.Scan(&obj.Id, &obj.Name, &obj.Sku, &obj.Stock, &obj.Price, &obj.Discount, &obj.Image, &obj.CategoryId)
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

func FindProductsById(id int) (Response, error) {
	var obj Products2
	var res Response

	con := db.CreateCon()

	sqlStatement := "SELECT id, name, sku, stock, price, discount, image, category_id FROM products WHERE id = ?"

	rows, err := con.Query(sqlStatement, id)
	if err != nil {
		return res, err
	}
	defer rows.Close()

	for rows.Next() {
		err = rows.Scan(&obj.Id, &obj.Name, &obj.Sku, &obj.Stock, &obj.Price, &obj.Discount, &obj.CategoryId)
		if err != nil {
			return res, err
		}
	}

	res.Success = true
	res.Message = "Success"
	res.Data = obj

	return res, nil
}

func StoreProducts(name string, image string, price int64, stock int64) (Response, error) {
	var res Response
	var obj Products

	con := db.CreateCon()

	sqlStatement := "INSERT products (name, price, stock, image, created_at, updated_at) VALUES (?, ?, ?, now(), nw());"

	stmt, err := con.Prepare(sqlStatement)
	if err != nil {
		return res, err
	}

	result, err := stmt.Exec(name, price, stock, image)
	if err != nil {
		return res, err
	}

	lastInsertedId, err := result.LastInsertId()
	if err != nil {
		return res, err
	}

	sqlStatement2 := "SELECT * FROM products WHERE id = ?;"
	rows, err := con.Query(sqlStatement2, lastInsertedId)

	defer rows.Close()

	if err != nil {
		return res, err
	}

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

func UpdateProducts(id int, name string, tipe string, logo string) (Response, error) {
	var res Response

	con := db.CreateCon()

	sqlStatement := "UPDATE products SET name= ?, type= ?, logo= ?, updated_at = now() WHERE id = ?;"

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

func DeleteProducts(id int) (Response, error) {
	var res Response

	con := db.CreateCon()

	sqlStatement := "DELETE FROM products WHERE id = ?;"

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
