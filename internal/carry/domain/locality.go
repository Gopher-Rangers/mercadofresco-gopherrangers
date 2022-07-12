package domain

type Locality struct {
	ID    int    `json:"locality_id"`
	Name  string `json:"locality_name"`
	Count int    `json:"carries_count"`
}
