package buyer

const (
	sqlGetAll = "SELECT * FROM `mercado-fresco`.`buyers`"

	sqlGetById = "SELECT * FROM `mercado-fresco`.`buyers` where id=?"

	sqlStore = "INSERT INTO `mercado-fresco`.`buyers` (`card_number_id`, `first_name`, `last_name`) VALUES (?, ?, ?)"

	sqlLastID = "SELECT MAX(id) as last_id FROM `mercado-fresco`.`buyers`"

	sqlUpdate = "UPDATE `mercado-fresco`.`buyers` SET card_number_id=?, first_name=?, last_name=? WHERE id=?"

	sqlUpdateName = "UPDATE `mercado-fresco`.`buyers` SET first_name=? WHERE id=?"

	sqlDelete = "DELETE FROM `mercado-fresco`.`buyers` WHERE id=?"
)
