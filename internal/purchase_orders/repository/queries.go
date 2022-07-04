package repository

const (
	sqlGetById = "SELECT * FROM `mercado-fresco`.`purchase_orders` where id=?"

	sqlCreate = "INSERT INTO `mercado-fresco`.`purchase_orders` (`order_number`, `order_date`, `tracking_code`, `buyer_id`, `product_record_id`, `order_status_id`) VALUES (?, ?, ?, ?, ?, ?)"

	sqlExistsOrderNumber = "SELECT 1 FROM `mercado-fresco`.`purchase_orders` WHERE `purchase_orders`.`order_number` = ?"
)
