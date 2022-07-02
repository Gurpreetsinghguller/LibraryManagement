package types

type Book struct {
	Name     string  `json:"name"`
	Author   string  `json:"author"`
	Category string  `json:"category"`
	Country  string  `json:"country"`
	Price    float32 `json:"price"`
}
