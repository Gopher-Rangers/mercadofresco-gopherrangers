package section

const (
	sqlGetAll = "SELECT * FROM `mercado-fresco`.`section`"

	sqlGetById = "SELECT * FROM `mercado-fresco`.`section` where id=?"

	sqlStore = "INSERT INTO `mercado-fresco`.`section` (`section_number`, `current_temperature`, `minimum_temperature`, `current_capacity`, `minimum_capacity`, `maximum_capacity`, `warehouse_id`, `product_type_id`) VALUES (?, ?, ?, ?, ?, ?, ?, ?)"

	sqlUpdateSecID = "UPDATE `mercado-fresco`.`section` SET section_number=? WHERE id=?"

	sqlDelete = "DELETE FROM `mercado-fresco`.`section` WHERE id=?"
)
