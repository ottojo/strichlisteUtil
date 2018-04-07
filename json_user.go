package strichlisteUtil

type User struct {
	ID              int           `json:"id"`
	Name            string        `json:"name"`
	MailAddress     string        `json:"mailAddress"`
	Balance         float64       `json:"balance"`
	LastTransaction string        `json:"lastTransaction"`
	Transactions    []Transaction `json:"transactions"`
}
