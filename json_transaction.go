package strichlisteUtil

type Transaction struct {
	ID         int     `json:"id"`
	UserID     int     `json:"userId"`
	CreateDate string  `json:"createDate"`
	Value      float64 `json:"value"`
}
