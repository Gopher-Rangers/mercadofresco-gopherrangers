package section

const (
	sqlGetAll = "SELECT * FROM section"

	sqlGetById = "SELECT * FROM section WHERE id=?"

	sqlStore = "INSERT INTO section (`section_number`, `current_temperature`, `minimum_temperature`, `current_capacity`, `minimum_capacity`, `maximum_capacity`, `warehouse_id`, `product_type_id`) VALUES (?, ?, ?, ?, ?, ?, ?, ?)"

	sqlUpdateSecID = "UPDATE section SET section_number=? WHERE id=?"

	sqlDelete = "DELETE FROM section WHERE id=?"
)
