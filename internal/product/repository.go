package products

type Product struct {
	ID int `json:"id"`
	ProductCode string `json:"product_code"`
	Description string `json:"description"`
	Width float64 `json:"width"`
	Height float64 `json:"height"`
	Length float64 `json:"length"`
	NetWeight float64 `json:"net_weight"`
	ExpirationRate string `json:"expiration_rate"`
	RecommendedFreezingTemperature float64 `json:"recommended_freezing_temperature"`
	FreezingRate float64 `json:"freezing_rate"`
	ProductTypeTd int `json:"product_type_id"`
	SellerId int `json:"seller_id"`
}
