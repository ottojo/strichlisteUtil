package strichlisteUtil

type usersResponse struct {
	OverallCount int         `json:"overallCount"`
	Limit        interface{} `json:"limit"`
	Offset       interface{} `json:"offset"`
	Entries      []User      `json:"entries"`
}
