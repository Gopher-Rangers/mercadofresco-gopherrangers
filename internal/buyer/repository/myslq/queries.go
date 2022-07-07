package myslq

const (
	SqlGetAll = "SELECT * FROM `mercado-fresco`.`buyers`"

	SqlGetById = "SELECT * FROM `mercado-fresco`.`buyers` where id=?"

	SqlBuyerWithOrdersById = "SELECT buyers.*, COUNT(purchase_orders.id) as purchase_orders_count\n  " +
		"FROM `mercado-fresco`.buyers \n  " +
		"JOIN `mercado-fresco`.purchase_orders \n    " +
		"ON purchase_orders.buyer_id = buyers.id\n" +
		"WHERE buyers.id = ?\n" +
		"GROUP BY buyers.id "

	SqlBuyersWithOrders = "SELECT buyers.*, COUNT(purchase_orders.id) as purchase_orders_count\n  " +
		"FROM `mercado-fresco`.buyers \n  " +
		"JOIN `mercado-fresco`.purchase_orders \n    " +
		"ON purchase_orders.buyer_id = buyers.id\n" +
		"GROUP BY buyers.id "

	SqlStore = "INSERT INTO `mercado-fresco`.`buyers` (`card_number_id`, `first_name`, `last_name`) VALUES (?, ?, ?)"

	SqlUpdate = "UPDATE `mercado-fresco`.`buyers` SET card_number_id=?, first_name=?, last_name=? WHERE id=?"

	SqlDelete = "DELETE FROM `mercado-fresco`.`buyers` WHERE id=?"
)
