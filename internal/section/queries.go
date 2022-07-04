package section

const (
	SqlGetAll = "SELECT * FROM section"

	SqlGetById = "SELECT * FROM section WHERE id=?"

	SqlStore = "INSERT INTO section (`section_number`, `current_temperature`, `minimum_temperature`, `current_capacity`, `minimum_capacity`, `maximum_capacity`, `warehouse_id`, `product_type_id`) VALUES (?, ?, ?, ?, ?, ?, ?, ?)"

	SqlUpdateSecID = "UPDATE section SET section_number=? WHERE id=?"

	SqlDelete = "DELETE FROM section WHERE id=?"
)
