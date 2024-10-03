package entity

type Vacancy struct {
	ID          int64  `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Updated_At  string `json:"updated_At"`
	Created_At  string `json:"created_At"`
}
