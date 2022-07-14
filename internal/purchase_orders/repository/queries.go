package repository

const (
	SqlGetById = "SELECT * FROM `mercado-fresco`.`purchase_orders` where id=?"

	SqlCreate = "INSERT INTO `mercado-fresco`.`purchase_orders` (`order_number`, `order_date`, `tracking_code`, `buyer_id`, `product_record_id`, `order_status_id`) VALUES (?, ?, ?, ?, ?, ?)"

	SqlOrderNumber = "SELECT order_number FROM `mercado-fresco`.`purchase_orders` WHERE `purchase_orders`.`order_number` = ?"
)
