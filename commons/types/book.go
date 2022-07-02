package types

type Book struct {
	Name     string  `json:"name" validate:"required"`
	Author   string  `json:"author" validate:"required"`
	Category string  `json:"category" validate:"required"`
	Country  string  `json:"country" validate:"required"`
	Price    float32 `json:"price" validate:"required"`
}
