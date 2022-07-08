package domain

type Carry struct {
	ID         int    `json:"id"`
	Cid        string `json:"cid" binding:"required"`
	Name       string `json:"name"`
	Address    string `json:"address"`
	Telephone  string `json:"telephone"`
	LocalityID int    `json:"locality_id"`
}
