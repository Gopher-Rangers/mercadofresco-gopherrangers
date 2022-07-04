package productbatch

const (
	sqlReportBatch = "SELECT * FROM section WHERE id=?"

	sqlCreateBatch = "INSERT INTO section (`section_number`, `current_temperature`, `minimum_temperature`, `current_capacity`, `minimum_capacity`, `maximum_capacity`, `warehouse_id`, `product_type_id`) VALUES (?, ?, ?, ?, ?, ?, ?, ?)"
)
