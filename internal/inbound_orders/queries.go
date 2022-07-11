package inboundorders

const (
	SqlCountByEmployee = "SELECT COUNT(*) FROM inbound_orders WHERE employee_id=?;"

	SqlCreate = "INSERT INTO inbound_orders (`order_date`, `order_number`, `employee_id`, `product_batch_id`, `warehouse_id`) VALUES (?, ?, ?, ?, ?)"
)
