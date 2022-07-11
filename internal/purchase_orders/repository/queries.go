package repository

const (
	SqlGetById = "SELECT * FROM `mercado-fresco`.`purchase_orders` where id=?"

	SqlCreate = "INSERT INTO `mercado-fresco`.`purchase_orders` (`order_number`, `order_date`, `tracking_code`, `buyer_id`, `product_record_id`, `order_status_id`) VALUES (?, ?, ?, ?, ?, ?)"

	SqlExistsOrderNumber = "SELECT 1 FROM `mercado-fresco`.`purchase_orders` WHERE `purchase_orders`.`order_number` = ?"
)
