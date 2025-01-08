package repository

import (
	"context"
	"fmt"

	"github.com/spf13/viper"
	"gorm.io/gorm"

	"face-be-service/common/constants"

	"face-be-service/common/database/entities"
	"face-be-service/common/utils"
)

type EmployeeRepository struct {
	db *gorm.DB
}

func NewEmployeeRepository(db *gorm.DB) *EmployeeRepository {
	return &EmployeeRepository{db: db}
}

func (ar *EmployeeRepository) Create(ctx context.Context, record *entities.Employee) error {
	return ar.db.Table("employees").WithContext(ctx).Create(&record).Error
}

func (ar *EmployeeRepository) GetTopByDistanceType(ctx context.Context, distanceType string,
	embedding []float32) (*entities.Employee, error) {
	var record *entities.Employee

	vector, err := utils.ConvertArrayFloat32(embedding)
	if err != nil {
		return nil, err
	}

	query := fmt.Sprintf(`
		SELECT id, employee_id, employee_name, image_path, embedding %v '%v' as score
		FROM employees
		WHERE embedding %v '%v' <= %v
		ORDER BY score ASC
		LIMIT 1`, distanceType, vector, distanceType, vector, viper.GetFloat64(constants.Threshold))

	err = ar.db.Raw(query).Scan(&record).Error
	if err != nil {
		return nil, err
	}

	return record, nil
}
