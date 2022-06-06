package seller

type Seller struct {
	Id          int    `json:"id"`
	Cid         int    `json:"company_id"`
	CompanyName string `json:"company_name"`
	Address     string `json:"address"`
	Telephone   string `json:"telephone"`
}
