package myslq

const (
	SqlGetAll = "SELECT * FROM `mercado-fresco`.`buyers`"

	SqlGetById = "SELECT * FROM `mercado-fresco`.`buyers` where id=?"

	SqlCountOrdersByBuyerId = "SELECT COUNT(*) FROM `mercado-fresco`.`purchase_orders` where buyer_id=?"

	SqlStore = "INSERT INTO `mercado-fresco`.`buyers` (`card_number_id`, `first_name`, `last_name`) VALUES (?, ?, ?)"

	sqlLastID = "SELECT MAX(id) as last_id FROM `mercado-fresco`.`buyers`"

	SqlUpdate = "UPDATE `mercado-fresco`.`buyers` SET card_number_id=?, first_name=?, last_name=? WHERE id=?"

	sqlUpdateName = "UPDATE `mercado-fresco`.`buyers` SET first_name=? WHERE id=?"

	SqlDelete = "DELETE FROM `mercado-fresco`.`buyers` WHERE id=?"
)
