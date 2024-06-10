package common

// Parameter request for page with pagination or search or filter
type Parameter struct {
	Cursor    string            `json:"cursor"`
	Direction string            `json:"direction"`
	Search    string            `json:"search"`
	Filter    map[string]string `json:"filter"`
}
