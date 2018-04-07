package strichlisteUtil

type transactionsResponse struct {
	OverallCount int         `json:"overallCount"`
	Limit        interface{} `json:"limit"`
	Offset       interface{} `json:"offset"`
	Entries      []Transaction`json:"entries"`
}
