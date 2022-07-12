package domain

type Warehouse struct {
	ID            int    `json:"id"`
	WarehouseCode string `json:"warehouse_code" binding:"required"`
	Address       string `json:"address"`
	Telephone     string `json:"telephone"`
	LocalityID    int    `json:"locality_id"`
}
