package models

type Todo struct {
	ID          int    `json:"id" gorm:"primary_key"`
	Title       string `json:"title" gorm:"not null"`
	Description string `json:"description" gorm:"null"`
	Status      bool   `json:"status" gorm:"not null"`
}

type CreateTodo struct {
	Title       string `json:"title" binding:"required"`
	Description string `json:"description"`
	Status      bool   `json:"status" binding:"required"`
}

type UpdateTodo struct {
	ID     int  `json:"id" binding:"required"`
	Status bool `json:"status" binding:"required"`
}
