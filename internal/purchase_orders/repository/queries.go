package repository

const (
	SqlGetById = "SELECT * FROM purchase_orders where id=?"

	SqlCreate = "INSERT INTO purchase_orders (`order_number`, `order_date`, `tracking_code`, `buyer_id`, `product_record_id`, `order_status_id`) VALUES (?, ?, ?, ?, ?, ?)"

	SqlOrderNumber = "SELECT order_number FROM purchase_orders where order_number = ?"
)
