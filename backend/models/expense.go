// package models

// type Expense struct {
// 	ID       int     `json:"id"`
// 	Name     string  `json:"name"`
// 	Amount   float64 `json:"amount"`
// 	Category string  `json:"category"`
// 	Date     string  `json:"date"`
// }

package models

type Expense struct {
	ID       int     `json:"id"`
	Name     string  `json:"name"`
	Amount   float64 `json:"amount"`
	Category string  `json:"category"`
	Date     string  `json:"date"` // Format: "YYYY-MM-DD"
}
