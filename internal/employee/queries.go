package employee

const (
	SqlGetAll = "SELECT * FROM employees"

	SqlGetById = "SELECT * FROM employees WHERE id=?"

	SqlCreate = "INSERT INTO employees (`card_number_id`, `first_name`, `last_name`, `warehouse_id`) VALUES (?, ?, ?, ?)"

	SqlUpdateFirstName = "UPDATE employees SET first_name=? WHERE id=?"

	SqlDelete = "DELETE FROM employees WHERE id=?"
)
