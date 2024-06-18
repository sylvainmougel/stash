package data

type Entry struct {
	ObjectID string  `json:"objectID"`
	Date     string  `json:"date"`
	Total    float64 `json:"total"`
	Epoch    int64   `json:"epoch"`
}
