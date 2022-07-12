package domain

type Warehouse struct {
	ID             int    `json:"id"`
	WarehouseCode  string `json:"warehouse_code" binding:"required"`
	Address        string `json:"address"`
	Telephone      string `json:"telephone"`
	MinCapacity    int    `json:"minimum_capacity"`
	MinTemperature int    `json:"minimum_temperature"`
}
