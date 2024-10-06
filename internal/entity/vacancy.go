package entity

type Vacancy struct {
	ID          int64  `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	UpdatedAt   string `json:"updated_At"`
	CreatedAt   string `json:"created_At"`
}
