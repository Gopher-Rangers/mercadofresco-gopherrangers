package productbatch

const (
	sqlReportBatchAll = "SELECT a.section_id, b.section_number, COUNT(*) FROM product_batches as a INNER JOIN section as b ON a.section_id = b.id GROUP BY section_id"

	sqlReportBatchByID = "SELECT a.section_id, b.section_number, COUNT(*) FROM product_batches as a INNER JOIN section as b ON a.section_id = b.id WHERE a.section_id = ? GROUP BY section_id"

	sqlCreateBatch = "INSERT INTO product_batches (batch_number, current_quantity, current_temperature, due_date, initial_quantity, manufacturing_date, manufacturing_hour, minimum_temperature, product_id, section_id) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?)"
)
