package repository

import (
	"context"

	"face-be-service/common/database/entities"
)

//go:generate mockgen -package=repository -destination=iemployee_mock.go -source=iemployee.go

type IEmployeeRepositoryInterface interface {
	Create(ctx context.Context, record *entities.Employee) error
	GetTopByDistanceType(ctx context.Context, distanceType string, embedding []float32) (*entities.Employee, error)
}
