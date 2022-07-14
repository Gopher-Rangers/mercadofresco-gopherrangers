package myslq

const (
	SqlGetAll = "SELECT * FROM buyers"

	SqlGetById = "SELECT * FROM buyers where id=?"

	SqlBuyerWithOrdersById = "SELECT buyers.*, COUNT(purchase_orders.id) as purchase_orders_count\n  " +
		"FROM buyers \n  " +
		"LEFT JOIN purchase_orders \n    " +
		"ON purchase_orders.buyer_id = buyers.id\n" +
		"WHERE buyers.id = ?\n" +
		"GROUP BY buyers.id "

	SqlBuyersWithOrders = "SELECT buyers.*, COUNT(purchase_orders.id) as purchase_orders_count\n  " +
		"FROM buyers \n  " +
		"LEFT JOIN purchase_orders \n    " +
		"ON purchase_orders.buyer_id = buyers.id\n" +
		"GROUP BY buyers.id "

	SqlStore = "INSERT INTO buyers (`card_number_id`, `first_name`, `last_name`) VALUES (?, ?, ?)"

	SqlUpdate = "UPDATE buyers SET card_number_id=?, first_name=?, last_name=? WHERE id=?"

	SqlDelete = "DELETE FROM buyers WHERE id=?"

	SqlUniqueCardNumberId = "SELECT id FROM buyers where id != ? and card_number_id = ?"
)
