package model

type Epic struct {
	Model

	ProjectID   int64  `json:"project_id"`
	SprintID    int64  `json:"sprint_id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Status      string `json:"status"`
	AssignTo    int64  `json:"assign_to"`
	CreatedBy   int64  `json:"created_by"`
	AssignUser  User   `gorm:"foreignKey:AssignTo"`
	CreatedUser User   `gorm:"foreignKey:CreatedBy"`
}
