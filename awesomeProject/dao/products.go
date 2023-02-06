package dao

import (
	"main/model"
)

// SearchProductsByID 依据商品ID检索商品信息//
func SearchProductsByID(ID int) (p model.Products, err error) {
	row := DB.QueryRow("select P_id,P_name,P_price,P_sales,P_comment,P_message from products where P_id=?", ID)
	if err = row.Err(); row.Err() != nil {
		return
	}
	err = row.Scan(&p.ProductsID, &p.ProductsName, &p.ProductsPrice, &p.ProductsSales, &p.ProductsComment, &p.ProductsMessage)
	return
}

func SearchProductsByID2(ID int) (s model.Cart, err error) {
	row := DB.QueryRow("select u_id from shopping_cart where p_id=?", ID)
	if err = row.Err(); row.Err() != nil {
		return
	}
	err = row.Scan(&s.UserID)
	return
}

func InsertInfo(s model.Cart) (err error) {
	_, err = DB.Exec("insert into shopping_cart(u_id,p_id,status) values (?,?,?)", s.UserID, s.ProductsID, s.Status)
	if err != nil {
		return
	}
	return
}

func DeleteCart(s model.Cart) (err error) {
	_, err = DB.Exec("delete from shopping_cart where p_id=?", s.ProductsID)
	if err != nil {
		return
	}
	//_, err = result.RowsAffected()
	//if err != nil {
	//	log.Printf("err:%s", err)
	//	return
	//}
	//_, err = result.LastInsertId()
	//if err != nil {
	//	log.Printf("err:%s", err)
	//	return
	//}
	return
}

func AlterStatus(s model.Cart) (err error) {
	_, err = DB.Exec("update shopping_cart set status=? where u_id=? and p_id=? ", s.Status, s.UserID, s.ProductsID)
	if err != nil {
		return
	}
	return
}

func SelectInfoByID(id int, i int) (s model.Cart, err error) {
	row := DB.QueryRow("select u_id,p_id,status from shopping_cart where u_id=? and p_id=? ", id, i)
	if err = row.Err(); row.Err() != nil {
		return
	}
	err = row.Scan(&s.UserID, &s.ProductsID, &s.Status)
	return
}

func SelectStatus(id int) (s model.Cart, err error) {
	row := DB.QueryRow("select status from shoppin_cart where u_id=?", id)
	if err = row.Err(); row.Err() != nil {
		return
	}
	err = row.Scan(&s.Status)
	return
}
