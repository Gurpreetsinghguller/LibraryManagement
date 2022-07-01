package types

type Book struct {
	ID       string  `json:"id"`
	Name     string  `json:"name"`
	Author   string  `json:"author"`
	Category string  `json:"category"`
	Country  string  `json:"country"`
	Price    float32 `json:"price"`
}
