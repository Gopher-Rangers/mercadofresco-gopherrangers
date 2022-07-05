package locality

import "github.com/Gopher-Rangers/mercadofresco-gopherrangers/internal/seller"

type Locality struct {
	Id           int             `json:"id"`
	LocalityName string          `json:"locality_name"`
	ProvinceName string          `json:"province_name"`
	CountryName  string          `json:"country_name"`
	Seller       []seller.Seller `json:"seller_id,omitempty"`
}

type ReportSeller struct {
	LocalityID   int    `json:"locality_id"`
	LocalityName string `json:"locality_name"`
	SellersCount int    `json:"sellers_count"`
}
