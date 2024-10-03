package model

type VacancyRequestCreate struct {
	Title       string `json:"title"`
	Description string `json:"description"`
}

type VacancyRequestUpdate struct {
	Title       string `json:"title"`
	Description string `json:"description"`
}
