package dto

type CreateBookRequest struct {
	Name        string `json:"name" validate:"required"`
	Stock       int    `json:"stock" validate:"required"`
	Description string `json:"description"`
}

type UpdateBookRequest struct {
	Name        *string `json:"name"`
	Stock       *int    `json:"stock"`
	Description *string `json:"description"`
}
