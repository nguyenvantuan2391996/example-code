package entities

import (
	"time"

	"github.com/pgvector/pgvector-go"
	"gorm.io/gorm"
)

type Employee struct {
	CreatedAt    *time.Time      `json:"created_at"`
	UpdatedAt    *time.Time      `json:"updated_at"`
	DeletedAt    gorm.DeletedAt  `json:"deleted_at"`
	ImagePath    string          `json:"image_path"`
	EmployeeName string          `json:"employee_name"`
	EmployeeID   string          `json:"employee_id"`
	Embedding    pgvector.Vector `json:"embedding" gorm:"type:vector(512)"`
	ID           int64           `json:"id"`
}
